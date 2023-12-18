package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	_ "github.com/go-sql-driver/mysql"
)

// CollectData connects to the RDS instance and CloudWatch, then collects and prints metrics
func CollectData() {
	// Connect to the database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbEndpoint, dbName))
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	// Example of direct database query
	var threadsConnected string
	err = db.QueryRow("SHOW STATUS LIKE 'Threads_connected'").Scan(&threadsConnected)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	fmt.Println("Threads Connected:", threadsConnected)

	// Connect to CloudWatch
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // Set your AWS region
	})
	if err != nil {
		fmt.Println("Error creating AWS session:", err)
		return
	}

	cw := cloudwatch.New(sess)
	getAndPrintMetrics(cw, "your_rds_instance_identifier") // Replace with your RDS instance identifier
}

// getAndPrintMetrics fetches metrics from CloudWatch and prints them
func getAndPrintMetrics(cw cloudwatchiface.CloudWatchAPI, instanceIdentifier string) {
	// Define the metric queries here
	// Example: CPUUtilization
	cpuMetric := &cloudwatch.MetricDataQuery{
		Id: aws.String("cpuUtilization"),
		MetricStat: &cloudwatch.MetricStat{
			Metric: &cloudwatch.Metric{
				Namespace:  aws.String("AWS/RDS"),
				MetricName: aws.String("CPUUtilization"),
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("DBInstanceIdentifier"),
						Value: aws.String(instanceIdentifier),
					},
				},
			},
			Period: aws.Int64(300), // in seconds
			Stat:   aws.String("Average"),
		},
	}

	// Fetch the metrics
	result, err := cw.GetMetricData(&cloudwatch.GetMetricDataInput{
		StartTime:         aws.Time(time.Now().Add(-30 * time.Minute)), // 30 minutes ago
		EndTime:           aws.Time(time.Now()),
		MetricDataQueries: []*cloudwatch.MetricDataQuery{cpuMetric},
	})
	if err != nil {
		fmt.Println("Error fetching metric data:", err)
		return
	}

	// Print the metrics
	for _, r := range result.MetricDataResults {
		fmt.Printf("%s: %v\n", *r.Label, r.Values)
	}
}
