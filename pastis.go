package pastis

import "github.com/aws/aws-lambda-go/lambda"

func Start(engine *Engine){
	lambda.Start(engine.Run)
}
