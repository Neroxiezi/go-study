package facade

import "fmt"

// 外观模式

type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

type AModuleAPI interface {
	TestA() string
}

type BModuleAPI interface {
	TestB() string
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type aModuleImpl struct{}
type bModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A-PFinal"
}
func (*bModuleImpl) TestB() string {
	return "B-PFinal"
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}
