//go:generate mockgen -source=interface.go -destination=iMock.go -package=tool
package tool

type myInterface interface {
	Int() int
}
