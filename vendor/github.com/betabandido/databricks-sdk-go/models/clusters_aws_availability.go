/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type ClustersAwsAvailability string

// List of ClustersAwsAvailability
const (
	SPOT               ClustersAwsAvailability = "SPOT"
	ON_DEMAND          ClustersAwsAvailability = "ON_DEMAND"
	SPOT_WITH_FALLBACK ClustersAwsAvailability = "SPOT_WITH_FALLBACK"
)