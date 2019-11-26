package tabular

import (
  "log"
  "github.com/bmeg/grip/kvi"
  "github.com/bmeg/grip/kvi/boltdb"
)

type TabularIndex struct {
  kv kvi.KVInterface
}

func NewTablularIndex(path string) (*TabularIndex, error) {
  out := TabularIndex{}
  kv, err := boltdb.NewKVInterface(path, kvi.Options{})
  if err != nil {
    return nil, err
  }
  out.kv = kv
  return &out, nil
}

type TSVIndex struct {
  kv kvi.KVInterface
  path string
  pathID uint64
  indexName string
  indexCol int
  header []string
  lineReader *LineReader
}

func (t *TabularIndex) IndexTSV(path string, indexName string) *TSVIndex {
  o := TSVIndex{kv:t.kv, path:path, indexName:indexName}
  o.Init()
  return &o
}

func (t *TSVIndex) Init() error {
  SetPathValue(t.kv, t.path, t.pathID)

  hasHeader := false
  var err error
  t.lineReader, err = NewLineReader(t.path)
  if err != nil {
    return err
  }
  cparse := CSVParse{}
  count := uint64(0)
  for line := range t.lineReader.ReadLines() {
    row := cparse.Parse(string(line.Text))
    if !hasHeader {
      t.header = row
      hasHeader = true
      for i := range row {
        if t.indexName == row[i] {
          t.indexCol = i
        }
      }
    } else {
      SetIDLine(t.kv, t.pathID, row[t.indexCol], count)
      SetLineOffset(t.kv, t.pathID, count, line.Offset)
      count++
    }
  }
  log.Printf("Found %d rows", count)
  return nil
}


func (t *TSVIndex) GetLineNumber(id string) uint64 {
  return GetIDLine(t.kv, t.pathID, id)
}

func (t *TSVIndex) GetLineText(lineNum uint64) []byte {
  offset := GetLineOffset(t.kv, t.pathID, lineNum)
  //cparse := CSVParse{}
  log.Printf("LineOffset: %d", offset)
  return t.lineReader.SeekRead(offset)
}
