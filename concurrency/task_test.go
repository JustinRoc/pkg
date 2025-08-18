package concurrency

import (
	"context"
	"time"

	"errors"
	"fmt"
	"testing"
)

func TestConTask(t *testing.T) {
	ctx := context.Background()
	tasks := []Task[*FileContent]{
		func(ctx context.Context) (*FileContent, error) {
			fileContent, err := Donwload1(ctx, "file1")
			if err != nil {
				return nil, err
			}
			return fileContent, nil
		},
		func(ctx context.Context) (*FileContent, error) {
			fileContent, err := Donwload2(ctx, "file2")
			if err != nil {
				return nil, err
			}
			return fileContent, nil
		},
		func(ctx context.Context) (*FileContent, error) {
			fileContent, err := Donwload3(ctx, "file3")
			if err != nil {
				return nil, err
			}
			return fileContent, nil
		},
	}
	results, err := ConTasks(ctx, tasks, 2)
	if err != nil {
		fmt.Println("ERR", err)
	}
	for i, rst := range results {
		if rst == nil { // 这里的判断逻辑需要依据Task函数返回值的类型来判断，如果返回值是结构体，那这里就不能这么搞了
			fmt.Printf("rst %d is nil\n", i)
			continue
		}
		fmt.Printf("file%d content: %s\n", i, string(rst.data))
	}
}

type FileContent struct {
	data []byte
}

func Donwload1(ctx context.Context, file string) (*FileContent, error) {
	fmt.Println("download file1 ...")
	time.Sleep(time.Second * 2)
	return &FileContent{data: []byte("file1 content")}, nil
}

func Donwload2(ctx context.Context, file string) (*FileContent, error) {
	fmt.Println("download file2 ...")
	time.Sleep(time.Second * 2)
	// return nil, errors.New("file2 download failed")
	return &FileContent{data: []byte("file2 content")}, nil
}

func Donwload3(ctx context.Context, file string) (*FileContent, error) {
	fmt.Println("download file3 ...")
	time.Sleep(time.Second * 2)
	return nil, errors.New("file3 download failed")
}
