/**
 * @Author: Huyantian
 * @Date: 2021/3/7 上午11:43
 */

package flow_test

import (
	"flow"
	"flow/action"
	"flow/core"
	"flow/node"
	"testing"
)

func TestFlow(t *testing.T) {
	f := flow.NewFlow(
		node.NewApply("发起人", core.Handler{Name: "发起人"}),
		node.NewSignAudit("会签",
			node.NewNormalAudit("会签1", core.Handler{Name: "会签1"}),
			node.NewNormalAudit("会签2", core.Handler{Name: "会签2"}),
			node.NewParallel("并行",
				node.NewNormalAudit("并行1", core.Handler{Name: "并行1"}),
				node.NewNormalAudit("并行2", core.Handler{Name: "并行2"}),
			),
			node.NewSerial("串行",
				node.NewNormalAudit("串行", core.Handler{Name: "串行1"}),
				node.NewNormalAudit("串行", core.Handler{Name: "串行2"}),
				node.NewParallel("并行",
					node.NewNormalAudit("并行1", core.Handler{Name: "并行1"}),
					node.NewNormalAudit("并行2", core.Handler{Name: "并行2"}),
				),
			),
		),
		node.NewNormalAudit("审批2", core.Handler{Name: "审批2"}),
	)
	f.Start()
	if err := action.Submit(f); err != nil {
		t.Error(err)
	}

	if err := action.Pass(f, core.Handler{Name: "会签1"}); err != nil {
		t.Error(err)
	}
	if err := action.Pass(f, core.Handler{Name: "会签2"}); err != nil {
		t.Error(err)
	}
	if err := action.Pass(f, core.Handler{Name: "并行1"}); err != nil {
		t.Error(err)
	}
	if err := action.Pass(f, core.Handler{Name: "串行1"}); err != nil {
		t.Error(err)
	}
	if err := action.Pass(f, core.Handler{Name: "串行2"}); err != nil {
		t.Error(err)
	}
	if err := action.Pass(f, core.Handler{Name: "并行1"}); err != nil {
		t.Error(err)
	}
	if err := action.Pass(f, core.Handler{Name: "审批2"}); err != nil {
		t.Error(err)
	}
	if !f.IsEnd() {
		t.Error(f)
	}

}
