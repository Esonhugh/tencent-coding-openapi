package baseClient

import "github.com/guonaihong/gout"

// Do is Warp of dataflow.Gout 's dataflow
func (a *ApiClientBase) Do() error {
	return a.setAuthHeader().DataFlow.Do()
}

// request interface define for every request Object has this func
type request interface {
	SetAction() string
}

// Call is Functional func for call any openapi
func (a *ApiClientBase) CallAction(request request, response any) (err error) {
	request.SetAction()
	a.POST(CommonApiEndPoint).SetJSON(request).BindJSON(response)
	return a.Do()
}

// CallAction is Funcational func for call any openapi but with get quesy.
func (a *ApiClientBase) Call(request request, response any) (err error) {
	action := request.SetAction()
	a.POST(CommonApiEndPoint).
		SetQuery(gout.H{
			"Action": action,
		}).SetJSON(request).BindJSON(response)
	return a.Do()
}
