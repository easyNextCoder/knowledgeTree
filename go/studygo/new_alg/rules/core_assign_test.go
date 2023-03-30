package rules

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test_assembleWarmAssignResultNew(t *testing.T) {
	type args struct {
		roundNum        int
		warmPlayerIndex int
		warmPlayerId    int64
	}
	tests := []struct {
		name string
		args args
		want *AssignResultNew
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{
			roundNum:        1,
			warmPlayerIndex: 0,
			warmPlayerId:    9981,
		}},
		{name: "test1", args: args{
			roundNum:        2,
			warmPlayerIndex: 0,
			warmPlayerId:    9981,
		}},
		{name: "test1", args: args{
			roundNum:        3,
			warmPlayerIndex: 0,
			warmPlayerId:    9981,
		}},
		{name: "test1", args: args{
			roundNum:        4,
			warmPlayerIndex: 0,
			warmPlayerId:    9981,
		}},
		{name: "test1", args: args{
			roundNum:        5,
			warmPlayerIndex: 0,
			warmPlayerId:    9981,
		}},
		//第二组测试
		{name: "test2", args: args{
			roundNum:        1,
			warmPlayerIndex: 3,
			warmPlayerId:    9982,
		}},
		{name: "test2", args: args{
			roundNum:        2,
			warmPlayerIndex: 3,
			warmPlayerId:    9982,
		}},
		{name: "test2", args: args{
			roundNum:        3,
			warmPlayerIndex: 3,
			warmPlayerId:    9982,
		}},
		{name: "test2", args: args{
			roundNum:        4,
			warmPlayerIndex: 3,
			warmPlayerId:    9982,
		}},
		{name: "test2", args: args{
			roundNum:        5,
			warmPlayerIndex: 3,
			warmPlayerId:    9982,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			for i := 0; i < 1000; i++ {
				if got := assembleWarmAssignResultNew(tt.args.roundNum, tt.args.warmPlayerIndex, tt.args.warmPlayerId); !reflect.DeepEqual(got, tt.want) {

				}
			}
			end := time.Now()
			fmt.Println(end.Sub(start).Milliseconds())

		})
	}
}
