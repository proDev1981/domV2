package dom

import "strconv"

type Events struct{
	Type string `json:"type"`
	Value string `json:"value"`
	Ref string `json:"ref"`
}
/// getter target of event
func (ev *Events) GetTarget()(ele *Element){

	for i, e := range Dom{
		res,_ := strconv.Atoi(ev.Ref)
		if e.ref == res {
			Dom[i].Value = ev.Value
			ele = e
		}
	}
	if ele == nil { ele = &Element{}}
	return
}
