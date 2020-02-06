package go_blkchain_merkle_tree

import (
	"crypto/md5"
	"fmt"
)

type Data struct {
	s string
}

func (data *Data) CalculateHash() ([]byte, error) {
	h := md5.New()
	if _, err := h.Write([]byte(data.s)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func (data *Data) Equals(other Content) (bool, error) {
	return data.s == other.(*Data).s, nil
}

func ExampleMerkleTree_VerifyTree() {
	data1 := Data{
		s: `111`,
	}
	data2 := Data{
		s: `222`,
	}
	tree, err := NewMerkleTree([]Content{&data1, &data2})
	if err != nil {
		panic(err)
	}
	bool_, err := tree.VerifyTree()
	if err != nil {
		panic(err)
	}
	fmt.Println(bool_)

	// Output:
	// true
}

func ExampleMerkleTree_VerifyContent() {
	data1 := Data{
		s: `111`,
	}
	data2 := Data{
		s: `222`,
	}
	tree, err := NewMerkleTree([]Content{&data1, &data2})
	if err != nil {
		panic(err)
	}
	data3 := Data{
		s: `333`,
	}
	bool_, err := tree.VerifyContent(&data3)
	if err != nil {
		panic(err)
	}
	fmt.Println(bool_)

	// Output:
	// false
}
