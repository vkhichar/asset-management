package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
)

const EventResource = "/events"
const AssetMaintenanceEvent = "ASSET_MAINTENANCE_ACTIVITY"

type EventService interface {
	PostAssetEvent(context.Context, *domain.Asset) (string, error)
	PostCreateUserEvent(ctx context.Context, user *domain.User) (string, error)
	PostUpdateUserEvent(context.Context, *domain.User) (string, error)
	PostCreateAssetEvent(ctx context.Context, asset *domain.Asset) (string, error)
	PostMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (string, error)
	PostAssetMaintenanceActivityEvent(ctx context.Context, resBody *domain.MaintenanceActivity) (string, error)
}

type eventSvc struct {
	client *http.Client
}

func NewEventService() EventService {
	return &eventSvc{
		client: &http.Client{
			Timeout: time.Second * time.Duration(config.GetEventApiTimeout()),
		},
	}
}

func (e *eventSvc) PostCreateUserEvent(ctx context.Context, user *domain.User) (string, error) {

	object := contract.CreateUserEvent{
		EventType: "user",
		Data:      user,
	}
	postBody, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("Error while marshaling in event service: %s", err.Error())
		return "", err
	}
	responseBody := bytes.NewReader(postBody)

	req, err := http.NewRequest("POST", config.GetEventServiceUrl()+"/events", responseBody)

	if err != nil {
		fmt.Printf("event service: error while newrequest: %s", err.Error())
		return "", err
	}

	resultCh := make(chan []byte)
	var resp *http.Response
	var errPost, errIORead error
	var body []byte
	errCh := hystrix.Go("event_api", func() error {
		req.Header.Add("Content-type", "application/json")
		resp, errPost = e.client.Do(req)
		if errPost != nil {
			fmt.Printf("Event service: Error while creating Post request. Error: %s", errPost.Error())
			return errPost
		}
		body, errIORead = ioutil.ReadAll(resp.Body)

		if errIORead != nil {
			fmt.Printf("Event Service : error while converting into byte stream: %s", errIORead.Error())
			return errIORead
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Event Service: Event not created. %d", resp.StatusCode)
			return errors.New("Event not created")
		}

		resultCh <- body
		defer resp.Body.Close()
		return nil
	}, nil)

	select {
	case res := <-resultCh:
		fmt.Printf("\nsuccess to get response from sub-system: %s", string(res))
		var responseObj contract.CreateUserEventResponse
		errJsonUnmar := json.Unmarshal(body, &responseObj)
		if errJsonUnmar != nil {
			fmt.Printf("Event service: Error while json unmarshal. Error: %s", errJsonUnmar.Error())
			return "", errJsonUnmar
		}
		eventId := strconv.Itoa(responseObj.Id)
		return eventId, nil

	case err := <-errCh:
		fmt.Printf("\nfailed to get response from sub-system: %s", err.Error())
		return "", err
	}

	// client := http.Client{
	// 	Timeout: 3 * time.Second,
	// }

	// resp, err := client.Do(re)

	// if err != nil {
	// 	fmt.Printf("Error in event service while getting response in client.do: %s", err.Error())
	// 	return "", err
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Error in event service :%s", err.Error())

	// }

	// var respObj contract.CreateUserEventResponse

	// errJsonUnmar := json.Unmarshal(body, &respObj)

	// if errJsonUnmar != nil {
	// 	fmt.Printf("Event service: Error while json unmarshal. Error: %s", errJsonUnmar.Error())
	// 	return "", errJsonUnmar
	// }

	// eventId := strconv.Itoa(respObj.Id)

	// return eventId, nil

}

func (e *eventSvc) PostAssetEvent(ctx context.Context, asset *domain.Asset) (string, error) {
	object := contract.UpdateAssetEvent{}
	object.EvenType = "asset"
	object.Data = asset
	Revent, err := json.Marshal(object)
	if err != nil {
		fmt.Printf("Error while Marshaling %s", err.Error())
		return "", err
	}
	r := bytes.NewReader(Revent)
	req, errReq := http.NewRequest("POST", config.GetEventServiceUrl()+"/events", r)
	if errReq != nil {
		fmt.Printf("Error in http request %s", errReq.Error())
		return "", errReq
	}
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resultCh := make(chan []byte)
	var resp *http.Response
	var errPost, errIORead error
	var body []byte
	errCh := hystrix.Go("event_api", func() error {
		req.Header.Add("Content-type", "application/json")
		resp, errPost = client.Do(req)
		if errPost != nil {
			fmt.Printf("Event service: Error while creating Post request. Error: %s", errPost.Error())
			return errPost
		}
		body, errIORead = ioutil.ReadAll(resp.Body)

		if errIORead != nil {
			fmt.Printf("Event Service : error while converting into byte stream: %s", errIORead.Error())
			return errIORead
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Event Service: Event not created. %d", resp.StatusCode)
			return errors.New("Event not created")
		}

		resultCh <- body
		defer resp.Body.Close()
		return nil
	}, nil)

	select {
	case res := <-resultCh:
		fmt.Printf("\nsuccess to get response from sub-system: %s", string(res))
		var responseObj contract.AssetEventResponse
		errJsonUnmar := json.Unmarshal(body, &responseObj)
		if errJsonUnmar != nil {
			fmt.Printf("Event service: Error while json unmarshal. Error: %s", errJsonUnmar.Error())
			return "", errJsonUnmar
		}
		eventId := strconv.Itoa(responseObj.ID)
		return eventId, nil

	case err := <-errCh:
		fmt.Printf("\nfailed to get response from sub-system: %s", err.Error())
		return "", err
	}

	// 	resp, errPost := client.Do(rec)
	// 	if errPost != nil {
	// 		fmt.Printf("Event service: Error while sending Post request to event. Error: %s", errPost.Error())
	// 		return "", errPost
	// 	}

	// 	body, errReadAll := ioutil.ReadAll(resp.Body)
	// 	if errReadAll != nil {
	// 		fmt.Printf("Error While Performing ReadAll %s", errReadAll.Error())
	// 		return "", errReadAll
	// 	}
	// 	var responseObj contract.AssetEventResponse
	// 	errJsonMarshal := json.Unmarshal(body, &responseObj)
	// 	if errJsonMarshal != nil {
	// 		fmt.Printf("Event Service : Error While UnMarshaling :%s", errJsonMarshal.Error())
	// 		return "", errJsonMarshal
	// 	}
	// 	eventId := strconv.Itoa(responseObj.ID)
	// 	return eventId, nil
}

func (e *eventSvc) PostCreateAssetEvent(ctx context.Context, asset *domain.Asset) (string, error) {
	obj := contract.CreateAssetEvent{}
	obj.EventType = "asset"
	obj.Data = asset
	postBody, errJson := json.Marshal(obj)
	if errJson != nil {
		fmt.Printf("Event Service: Error while marshaling %s", errJson.Error())
		return "", errJson
	}

	responseBody := bytes.NewReader(postBody)
	req, errNewReq := http.NewRequest("POST", config.GetEventServiceUrl()+"/events", responseBody)
	if errNewReq != nil {
		fmt.Printf("Event Service: error during http request %s:", errNewReq.Error())
		return "", errNewReq
	}

	resultCh := make(chan []byte)
	var resp *http.Response
	var errPost, errIORead error
	var body []byte
	errCh := hystrix.Go("event_api", func() error {
		req.Header.Add("Content-type", "application/json")
		resp, errPost = e.client.Do(req)
		if errPost != nil {
			fmt.Printf("Event service: Error while creating Post request. Error: %s", errPost.Error())
			return errPost
		}
		body, errIORead = ioutil.ReadAll(resp.Body)

		if errIORead != nil {
			fmt.Printf("Event Service : error while converting into byte stream: %s", errIORead.Error())
			return errIORead
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Event Service: Event not created. %d", resp.StatusCode)
			return errors.New("Event not created")
		}

		resultCh <- body
		defer resp.Body.Close()
		return nil
	}, nil)

	select {
	case res := <-resultCh:
		fmt.Printf("\nsuccess to get response from sub-system: %s", string(res))
		var responseObj contract.CreateAssetEventResponse
		errJsonUnmar := json.Unmarshal(body, &responseObj)
		if errJsonUnmar != nil {
			fmt.Printf("Event service: Error while json unmarshal. Error: %s", errJsonUnmar.Error())
			return "", errJsonUnmar
		}
		eventId := strconv.Itoa(responseObj.ID)
		return eventId, nil

	case err := <-errCh:
		fmt.Printf("\nfailed to get response from sub-system: %s", err.Error())
		return "", err
	}

}

func (evSvc *eventSvc) PostUpdateUserEvent(ctx context.Context, user *domain.User) (string, error) {

	request := contract.UpdateUserEventRequest{}
	request.EventType = "user"
	request.Data = user
	reqEvent, errJson := json.Marshal(request)

	if errJson != nil {
		fmt.Printf("Event service: Error while json marshal. Error: %s", errJson.Error())
		return "", errJson
	}

	req, errNewReq := http.NewRequest("POST", config.GetEventServiceUrl()+EventResource, bytes.NewReader(reqEvent))
	if errNewReq != nil {
		fmt.Printf("Event service: Error while creating Post request. Error: %s", errNewReq.Error())
		return "", errNewReq
	}

	resultCh := make(chan []byte)
	var resp *http.Response
	var errPost, errIORead error
	var body []byte
	errCh := hystrix.Go("event_api", func() error {
		req.Header.Add("Content-type", "application/json")
		resp, errPost = evSvc.client.Do(req)
		if errPost != nil {
			fmt.Printf("Event service: Error while creating Post request. Error: %s", errPost.Error())
			return errPost
		}
		body, errIORead = ioutil.ReadAll(resp.Body)

		if errIORead != nil {
			fmt.Printf("Event Service : error while converting into byte stream: %s", errIORead.Error())
			return errIORead
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Event Service: Event not created. %d", resp.StatusCode)
			return errors.New("Event not created")
		}

		resultCh <- body
		defer resp.Body.Close()
		return nil
	}, nil)

	select {
	case res := <-resultCh:
		fmt.Printf("\nsuccess to get response from sub-system: %s", string(res))
		var responseObj contract.UpdateUserEventResponse
		errJsonUnmar := json.Unmarshal(body, &responseObj)
		if errJsonUnmar != nil {
			fmt.Printf("Event service: Error while json unmarshal. Error: %s", errJsonUnmar.Error())
			return "", errJsonUnmar
		}
		eventId := strconv.Itoa(responseObj.Id)
		return eventId, nil

	case err := <-errCh:
		fmt.Printf("\nfailed to get response from sub-system: %s", err.Error())
		return "", err
	}

	// req.Header.Add("Content-type", "application/json")
	// resp, errPost := evSvc.client.Do(req)
	// if errPost != nil {
	// 	fmt.Printf("Event service: Error while sending Post request to event. Error: %s", errPost.Error())
	// 	return "", errPost
	// }

	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Printf("Event Service: Event not created. %d", resp.StatusCode)
	// 	return "", errors.New("Event not created")
	// }

	// body, errBodyRead := ioutil.ReadAll(resp.Body)
	// defer resp.Body.Close()

	// if errBodyRead != nil {
	// 	fmt.Printf("Event service: Error while reading response body. Error: %s", errBodyRead.Error())
	// 	return "", errBodyRead
	// }
	// var responseObj contract.UpdateUserEventResponse

	// errJsonUnmar := json.Unmarshal(body, &responseObj)

	// if errJsonUnmar != nil {
	// 	fmt.Printf("Event service: Error while json unmarshal. Error: %s", errJsonUnmar.Error())
	// 	return "", errJsonUnmar
	// }

	// eventId := strconv.Itoa(responseObj.Id)

	// return eventId, nil
}

func (service *eventSvc) PostMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (string, error) {

	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	eventReqBody, _ := json.Marshal(contract.NewEventRequest(AssetMaintenanceEvent, reqBody))

	httpreq, err := http.NewRequest("POST", config.GetEventServiceUrl()+EventResource, bytes.NewBuffer(eventReqBody))
	// httpreq.Header.Add("Content-type", "application/json")
	// res, err := service.client.Do(httpreq)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }

	// if res.StatusCode != http.StatusOK {
	// 	fmt.Println("Failed to create event due to ", res.StatusCode)
	// 	return "", errors.New("Event not created")
	// }

	// resBody, err := ioutil.ReadAll(res.Body)
	// defer res.Body.Close()

	// if err != nil {
	// 	fmt.Println("Failed to read response\n", err)
	// 	return "", err
	// }

	// return string(resBody), err

	resultCh := make(chan []byte)
	var resp *http.Response
	var errPost, errIORead error
	var body []byte
	errCh := hystrix.Go("event_api", func() error {
		httpreq.Header.Add("Content-type", "application/json")
		resp, errPost = service.client.Do(httpreq)
		if errPost != nil {
			fmt.Printf("Event service: Error while creating Post request. Error: %s", errPost.Error())
			return errPost
		}
		body, errIORead = ioutil.ReadAll(resp.Body)

		if errIORead != nil {
			fmt.Printf("Event Service : error while converting into byte stream: %s", errIORead.Error())
			return errIORead
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Event Service: Event not created. %d", resp.StatusCode)
			return errors.New("Event not created")
		}

		resultCh <- body
		defer resp.Body.Close()
		return nil
	}, nil)

	select {
	case res := <-resultCh:
		fmt.Printf("\nsuccess to get response from sub-system: %s", string(res))
		return string(body), err

	case err := <-errCh:
		fmt.Printf("\nfailed to get response from sub-system: %s", err.Error())
		return "", err
	}
}

func (evSvc *eventSvc) PostAssetMaintenanceActivityEvent(ctx context.Context, resBody *domain.MaintenanceActivity) (string, error) {
	req := contract.CreateAssetMaintenanceEventRequest{}
	req.EventType = "maintenanceactivity"
	req.Data = resBody
	reqEvent, _ := json.Marshal(req)
	r := bytes.NewReader(reqEvent)

	reqst, err := http.NewRequest("POST", config.GetEventServiceUrl()+"/events", r)
	reqst.Header.Add("Content-type", "application/json")
	if err != nil {
		fmt.Printf("Event service request: error:%s", err.Error())
		return "", err
	}

	// resp, err := evSvc.client.Do(reqst)
	// if err != nil {
	// 	fmt.Printf("Event service error while getting response from client.Do: error:%s", err.Error())
	// 	return "", customerrors.ResponseTimeLimitExceeded
	// }

	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Println("Error while creating event")
	// 	return "", errors.New("Event not created")
	// }
	// body, errRead := ioutil.ReadAll(resp.Body)
	// if errRead != nil {
	// 	fmt.Printf("Event service read: error:%s", errRead.Error())
	// 	return "", errRead
	// }

	// errJsonUnmarshl := json.Unmarshal(body, &contract.CreateMaintenanceActivityEvent)
	// if errJsonUnmarshl != nil {
	// 	fmt.Printf("Event service error in unmarshaling: error:%s", errJsonUnmarshl.Error())
	// 	return "", errJsonUnmarshl
	// }

	resultCh := make(chan []byte)
	var resp *http.Response
	var errPost, errIORead error
	var body []byte
	errCh := hystrix.Go("event_api", func() error {
		resp.Header.Add("Content-type", "application/json")
		resp, errPost = evSvc.client.Do(reqst)
		if errPost != nil {
			fmt.Printf("Event service: Error while creating Post request. Error: %s", errPost.Error())
			return customerrors.ResponseTimeLimitExceeded
		}
		body, errIORead = ioutil.ReadAll(resp.Body)

		if errIORead != nil {
			fmt.Printf("Event Service : error while converting into byte stream: %s", errIORead.Error())
			return errIORead
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Event Service: Event not created. %d", resp.StatusCode)
			return errors.New("Event not created")
		}

		resultCh <- body
		defer resp.Body.Close()
		return nil
	}, nil)

	select {
	case res := <-resultCh:
		fmt.Printf("\nsuccess to get response from sub-system: %s", string(res))
		eventConvert := strconv.Itoa(contract.CreateMaintenanceActivityEvent.Id)
		return eventConvert, nil

	case err := <-errCh:
		fmt.Printf("\nfailed to get response from sub-system: %s", err.Error())
		return "", err
	}
}
