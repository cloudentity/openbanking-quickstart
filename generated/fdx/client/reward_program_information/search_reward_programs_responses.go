// Code generated by go-swagger; DO NOT EDIT.

package reward_program_information

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-quickstart/generated/fdx/models"
)

// SearchRewardProgramsReader is a Reader for the SearchRewardPrograms structure.
type SearchRewardProgramsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchRewardProgramsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchRewardProgramsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchRewardProgramsOK creates a SearchRewardProgramsOK with default headers values
func NewSearchRewardProgramsOK() *SearchRewardProgramsOK {
	return &SearchRewardProgramsOK{}
}

/* SearchRewardProgramsOK describes a response with status code 200, with default header values.

Data describing reward programs associated with accounts
*/
type SearchRewardProgramsOK struct {
	Payload *models.RewardProgramsentity
}

func (o *SearchRewardProgramsOK) Error() string {
	return fmt.Sprintf("[GET /reward-programs][%d] searchRewardProgramsOK  %+v", 200, o.Payload)
}
func (o *SearchRewardProgramsOK) GetPayload() *models.RewardProgramsentity {
	return o.Payload
}

func (o *SearchRewardProgramsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RewardProgramsentity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
