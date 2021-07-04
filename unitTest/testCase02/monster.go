package monster

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func (this *Monster) Store() bool {
	
	str, err := json.Marshal(this)
	if err != nil {
		fmt.Println("Marshall FAIL")
		return false
	}
	fileName := "monster.ser"

	err = ioutil.WriteFile(fileName,str,7555)
	if err != nil {
		fmt.Println("WriteFile FAIL")
		return false
	}

	return true

}

func (this *Monster) ReStore() bool {
	
	fileName := "./monster.ser"

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("ReadFile FAIL")
		return false
	}

	err = json.Unmarshal(content, this)
	if err != nil {
		
		fmt.Println("UnMarshall FAIL")
		return false
	}

	return true

}