package mongo

import (
	"context"
	"strconv"
	"strings"

	"github.com/bmeg/grip/gdbi"
	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"github.com/bmeg/grip/util"
	"google.golang.org/protobuf/types/known/structpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Processor stores the information for a mongo aggregation pipeline
type Processor struct {
	db              *Graph
	startCollection string
	query           mongo.Pipeline
	dataType        gdbi.DataType
	markTypes       map[string]gdbi.DataType
	aggTypes        map[string]*gripql.Aggregate
}

func getDataElement(result map[string]interface{}) *gdbi.DataElement {
	de := &gdbi.DataElement{}
	if x, ok := result["_id"]; ok {
		de.ID = x.(string)
	}
	if x, ok := result["label"]; ok {
		de.Label = x.(string)
	}
	if x, ok := result["data"]; ok {
		de.Data = x.(map[string]interface{})
	}
	if x, ok := result["to"]; ok {
		de.To = x.(string)
	}
	if x, ok := result["from"]; ok {
		de.From = x.(string)
	}
	return de
}

// Process runs the mongo aggregation pipeline
func (proc *Processor) Process(ctx context.Context, man gdbi.Manager, in gdbi.InPipe, out gdbi.OutPipe) context.Context {
	plog := log.WithFields(log.Fields{"query_id": util.UUID(), "query": proc.query, "query_collection": proc.startCollection})
	plog.Debug("Running Mongo Processor")

	go func() {
		defer close(out)

		initCol := proc.db.ar.client.Database(proc.db.ar.database).Collection(proc.startCollection)
		for t := range in {
			nResults := 0
			//plog.Infof("Running: %#v", proc.query)
			cursor, err := initCol.Aggregate(ctx, proc.query)
			if err != nil {
				plog.Errorf("Query Error (%s) : %s", proc.query, err)
				continue
			}
			//defer cursor.Close(context.TODO())
			result := map[string]interface{}{}
			for cursor.Next(ctx) {
				nResults++
				select {
				case <-ctx.Done():
					return
				default:
				}
				if nil != cursor.Decode(&result) {
					plog.Errorf("Result Error : %s", err)
					continue
				}
				switch proc.dataType {
				case gdbi.CountData:
					eo := &gdbi.Traveler{}
					if x, ok := result["count"]; ok {
						eo.Count = uint32(x.(int32))
					}
					out <- eo

				case gdbi.SelectionData:
					selections := map[string]*gdbi.DataElement{}
					if marks, ok := result["marks"]; ok {
						if marks, ok := marks.(map[string]interface{}); ok {
							for k, v := range marks {
								if v, ok := v.(map[string]interface{}); ok {
									de := getDataElement(v)
									selections[k] = de
								}
							}
						}
					}
					out <- &gdbi.Traveler{Selections: selections}

				case gdbi.AggregationData:

					for k, v := range result {
						buckets, ok := v.(bson.A)
						if !ok {
							plog.Errorf("Failed to convert Mongo aggregation result (%s): %+v", k, v)
							continue
						}

						var lastBucket float64
						for i, bucket := range buckets {
							bucket, ok := bucket.(map[string]interface{})
							if !ok {
								plog.Errorf("Failed to convert Mongo aggregation result bucket: %+v", bucket)
								continue
							}

							var term *structpb.Value
							switch proc.aggTypes[k].GetAggregation().(type) {
							case *gripql.Aggregate_Term:
								term, _ = structpb.NewValue(bucket["_id"])
							case *gripql.Aggregate_Histogram:
								term, _ = structpb.NewValue(bucket["_id"])
								curPos := bucket["_id"].(float64)
								stepSize := float64(proc.aggTypes[k].GetHistogram().Interval)
								if i != 0 {
									for nv := lastBucket + stepSize; nv < curPos; nv += stepSize {
										snv, _ := structpb.NewValue(nv)
										out <- &gdbi.Traveler{Aggregation: &gdbi.Aggregate{Name: k, Key: snv, Value: float64(0.0)}}
									}
								}
								lastBucket = curPos

							case *gripql.Aggregate_Percentile:
								bid := strings.Replace(bucket["_id"].(string), "_", ".", -1)
								f, err := strconv.ParseFloat(bid, 64)
								if err != nil {
									plog.Errorf("failed to parse percentile aggregation result key: %v", err)
									continue
								}
								term, _ = structpb.NewValue(f)
							default:
								plog.Errorf("unknown aggregation result type")
							}

							switch bucket["count"].(type) {
							case int:
								count := bucket["count"].(int)
								out <- &gdbi.Traveler{Aggregation: &gdbi.Aggregate{Name: k, Key: term, Value: float64(count)}}
							case int32:
								count := bucket["count"].(int32)
								out <- &gdbi.Traveler{Aggregation: &gdbi.Aggregate{Name: k, Key: term, Value: float64(count)}}
							case float64:
								count := bucket["count"].(float64)
								out <- &gdbi.Traveler{Aggregation: &gdbi.Aggregate{Name: k, Key: term, Value: float64(count)}}
							default:
								plog.Errorf("unexpected aggregation result type: %T", bucket["count"])
								continue
							}
						}
					}

				default:
					if marks, ok := result["marks"]; ok {
						if marks, ok := marks.(map[string]interface{}); ok {
							for k, v := range marks {
								if v, ok := v.(map[string]interface{}); ok {
									de := getDataElement(v)
									t = t.AddMark(k, de)
								}
							}
						}
					}

					de := &gdbi.DataElement{}
					if x, ok := result["_id"]; ok {
						de.ID = x.(string)
					}
					if x, ok := result["label"]; ok {
						de.Label = x.(string)
					}
					if x, ok := result["data"]; ok {
						de.Data = x.(map[string]interface{})
					}
					if x, ok := result["to"]; ok {
						de.To = x.(string)
					}
					if x, ok := result["from"]; ok {
						de.From = x.(string)
					}
					out <- t.AddCurrent(de)
				}
			}
			if err := cursor.Close(context.TODO()); err != nil {
				plog.WithFields(log.Fields{"error": err}).Error("MongoDb: iterating results")
				continue
			}
			if nResults == 0 && proc.dataType == gdbi.CountData {
				out <- &gdbi.Traveler{Count: 0}
			}
		}
		plog.Debug("Mongo Processor Finished")
	}()
	return ctx
}
