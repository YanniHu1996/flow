/**
 * @Author: Huyantian
 * @Date: 2021/3/7 上午11:45
 */

package action

import (
	"errors"
	"flow/core"
)

type Action func(*core.Flow)

func Submit(f *core.Flow) error {
	if f.IsEnd() {
		return errors.New("流程已结束")
	}
	f.FirstReady().Leave()
	return nil
}

func Pass(f *core.Flow, handler core.Handler) error {
	if f.IsEnd() {
		return errors.New("流程已结束")
	}
	var flag bool
	for _, node := range f.ReadyNodes {
		if node.Handler.Name == handler.Name {
			node.Leave()
			flag = true
		}
	}
	if !flag {
		return errors.New("错误的审批人")
	}
	return nil
}
