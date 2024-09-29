// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"crud/domain"
	"sync"
)

type FakeProductService struct {
	CreateProductStub        func(domain.Product) (domain.Product, error)
	createProductMutex       sync.RWMutex
	createProductArgsForCall []struct {
		arg1 domain.Product
	}
	createProductReturns struct {
		result1 domain.Product
		result2 error
	}
	createProductReturnsOnCall map[int]struct {
		result1 domain.Product
		result2 error
	}
	GetProductStub        func(string) (domain.Product, error)
	getProductMutex       sync.RWMutex
	getProductArgsForCall []struct {
		arg1 string
	}
	getProductReturns struct {
		result1 domain.Product
		result2 error
	}
	getProductReturnsOnCall map[int]struct {
		result1 domain.Product
		result2 error
	}
	GetProductsStub        func(int, int) ([]domain.Product, error)
	getProductsMutex       sync.RWMutex
	getProductsArgsForCall []struct {
		arg1 int
		arg2 int
	}
	getProductsReturns struct {
		result1 []domain.Product
		result2 error
	}
	getProductsReturnsOnCall map[int]struct {
		result1 []domain.Product
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProductService) CreateProduct(arg1 domain.Product) (domain.Product, error) {
	fake.createProductMutex.Lock()
	ret, specificReturn := fake.createProductReturnsOnCall[len(fake.createProductArgsForCall)]
	fake.createProductArgsForCall = append(fake.createProductArgsForCall, struct {
		arg1 domain.Product
	}{arg1})
	stub := fake.CreateProductStub
	fakeReturns := fake.createProductReturns
	fake.recordInvocation("CreateProduct", []interface{}{arg1})
	fake.createProductMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProductService) CreateProductCallCount() int {
	fake.createProductMutex.RLock()
	defer fake.createProductMutex.RUnlock()
	return len(fake.createProductArgsForCall)
}

func (fake *FakeProductService) CreateProductCalls(stub func(domain.Product) (domain.Product, error)) {
	fake.createProductMutex.Lock()
	defer fake.createProductMutex.Unlock()
	fake.CreateProductStub = stub
}

func (fake *FakeProductService) CreateProductArgsForCall(i int) domain.Product {
	fake.createProductMutex.RLock()
	defer fake.createProductMutex.RUnlock()
	argsForCall := fake.createProductArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeProductService) CreateProductReturns(result1 domain.Product, result2 error) {
	fake.createProductMutex.Lock()
	defer fake.createProductMutex.Unlock()
	fake.CreateProductStub = nil
	fake.createProductReturns = struct {
		result1 domain.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeProductService) CreateProductReturnsOnCall(i int, result1 domain.Product, result2 error) {
	fake.createProductMutex.Lock()
	defer fake.createProductMutex.Unlock()
	fake.CreateProductStub = nil
	if fake.createProductReturnsOnCall == nil {
		fake.createProductReturnsOnCall = make(map[int]struct {
			result1 domain.Product
			result2 error
		})
	}
	fake.createProductReturnsOnCall[i] = struct {
		result1 domain.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeProductService) GetProduct(arg1 string) (domain.Product, error) {
	fake.getProductMutex.Lock()
	ret, specificReturn := fake.getProductReturnsOnCall[len(fake.getProductArgsForCall)]
	fake.getProductArgsForCall = append(fake.getProductArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetProductStub
	fakeReturns := fake.getProductReturns
	fake.recordInvocation("GetProduct", []interface{}{arg1})
	fake.getProductMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProductService) GetProductCallCount() int {
	fake.getProductMutex.RLock()
	defer fake.getProductMutex.RUnlock()
	return len(fake.getProductArgsForCall)
}

func (fake *FakeProductService) GetProductCalls(stub func(string) (domain.Product, error)) {
	fake.getProductMutex.Lock()
	defer fake.getProductMutex.Unlock()
	fake.GetProductStub = stub
}

func (fake *FakeProductService) GetProductArgsForCall(i int) string {
	fake.getProductMutex.RLock()
	defer fake.getProductMutex.RUnlock()
	argsForCall := fake.getProductArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeProductService) GetProductReturns(result1 domain.Product, result2 error) {
	fake.getProductMutex.Lock()
	defer fake.getProductMutex.Unlock()
	fake.GetProductStub = nil
	fake.getProductReturns = struct {
		result1 domain.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeProductService) GetProductReturnsOnCall(i int, result1 domain.Product, result2 error) {
	fake.getProductMutex.Lock()
	defer fake.getProductMutex.Unlock()
	fake.GetProductStub = nil
	if fake.getProductReturnsOnCall == nil {
		fake.getProductReturnsOnCall = make(map[int]struct {
			result1 domain.Product
			result2 error
		})
	}
	fake.getProductReturnsOnCall[i] = struct {
		result1 domain.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeProductService) GetProducts(arg1 int, arg2 int) ([]domain.Product, error) {
	fake.getProductsMutex.Lock()
	ret, specificReturn := fake.getProductsReturnsOnCall[len(fake.getProductsArgsForCall)]
	fake.getProductsArgsForCall = append(fake.getProductsArgsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	stub := fake.GetProductsStub
	fakeReturns := fake.getProductsReturns
	fake.recordInvocation("GetProducts", []interface{}{arg1, arg2})
	fake.getProductsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProductService) GetProductsCallCount() int {
	fake.getProductsMutex.RLock()
	defer fake.getProductsMutex.RUnlock()
	return len(fake.getProductsArgsForCall)
}

func (fake *FakeProductService) GetProductsCalls(stub func(int, int) ([]domain.Product, error)) {
	fake.getProductsMutex.Lock()
	defer fake.getProductsMutex.Unlock()
	fake.GetProductsStub = stub
}

func (fake *FakeProductService) GetProductsArgsForCall(i int) (int, int) {
	fake.getProductsMutex.RLock()
	defer fake.getProductsMutex.RUnlock()
	argsForCall := fake.getProductsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProductService) GetProductsReturns(result1 []domain.Product, result2 error) {
	fake.getProductsMutex.Lock()
	defer fake.getProductsMutex.Unlock()
	fake.GetProductsStub = nil
	fake.getProductsReturns = struct {
		result1 []domain.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeProductService) GetProductsReturnsOnCall(i int, result1 []domain.Product, result2 error) {
	fake.getProductsMutex.Lock()
	defer fake.getProductsMutex.Unlock()
	fake.GetProductsStub = nil
	if fake.getProductsReturnsOnCall == nil {
		fake.getProductsReturnsOnCall = make(map[int]struct {
			result1 []domain.Product
			result2 error
		})
	}
	fake.getProductsReturnsOnCall[i] = struct {
		result1 []domain.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeProductService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createProductMutex.RLock()
	defer fake.createProductMutex.RUnlock()
	fake.getProductMutex.RLock()
	defer fake.getProductMutex.RUnlock()
	fake.getProductsMutex.RLock()
	defer fake.getProductsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProductService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ domain.ProductService = new(FakeProductService)
