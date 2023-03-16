package tests

import (
	"encoding/json"
	"testing"
)

func TestBuildMatrixOpQuery(t *testing.T) {
	m := MatrixOpQuery{}
	req := m.Build("getMatrix",10)
	t.Log(req)
}

func TestPopulateMatrixOpQueryResult(t *testing.T){
	data := `{"data":[[{"isEmptyCell":true}]]}`
	m := MatrixOpQuery{}
	err := json.Unmarshal([]byte(data),&m)
	if err != nil {
		t.Fatal(err)
	}
	if !m.Data[0][0].IsEmptyCell {
		t.Fail()
	}
	t.Logf("\n%#v",m)
	out,_ := json.MarshalIndent(m,"","\t")
	t.Logf("\n%s",string(out))
}