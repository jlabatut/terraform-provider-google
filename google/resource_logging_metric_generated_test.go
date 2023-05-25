// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccLoggingMetric_loggingMetricBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLoggingMetricDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingMetric_loggingMetricBasicExample(context),
			},
			{
				ResourceName:      "google_logging_metric.logging_metric",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccLoggingMetric_loggingMetricBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_logging_metric" "logging_metric" {
  name   = "tf-test-my-(custom)/metric%{random_suffix}"
  filter = "resource.type=gae_app AND severity>=ERROR"
  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "DISTRIBUTION"
    unit        = "1"
    labels {
      key         = "mass"
      value_type  = "STRING"
      description = "amount of matter"
    }
    labels {
      key         = "sku"
      value_type  = "INT64"
      description = "Identifying number for item"
    }
    display_name = "My metric"
  }
  value_extractor = "EXTRACT(jsonPayload.request)"
  label_extractors = {
    "mass" = "EXTRACT(jsonPayload.request)"
    "sku"  = "EXTRACT(jsonPayload.id)"
  }
  bucket_options {
    linear_buckets {
      num_finite_buckets = 3
      width              = 1
      offset             = 1
    }
  }
}
`, context)
}

func TestAccLoggingMetric_loggingMetricCounterBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLoggingMetricDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingMetric_loggingMetricCounterBasicExample(context),
			},
			{
				ResourceName:      "google_logging_metric.logging_metric",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccLoggingMetric_loggingMetricCounterBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_logging_metric" "logging_metric" {
  name   = "tf-test-my-(custom)/metric%{random_suffix}"
  filter = "resource.type=gae_app AND severity>=ERROR"
  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "INT64"
  }
}
`, context)
}

func TestAccLoggingMetric_loggingMetricCounterLabelsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLoggingMetricDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingMetric_loggingMetricCounterLabelsExample(context),
			},
			{
				ResourceName:      "google_logging_metric.logging_metric",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccLoggingMetric_loggingMetricCounterLabelsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_logging_metric" "logging_metric" {
  name   = "tf-test-my-(custom)/metric%{random_suffix}"
  filter = "resource.type=gae_app AND severity>=ERROR"
  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "INT64"
    labels {
      key         = "mass"
      value_type  = "STRING"
      description = "amount of matter"
    }
  }
  label_extractors = {
    "mass" = "EXTRACT(jsonPayload.request)"
  }
}
`, context)
}

func TestAccLoggingMetric_loggingMetricLoggingBucketExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLoggingMetricDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingMetric_loggingMetricLoggingBucketExample(context),
			},
			{
				ResourceName:      "google_logging_metric.logging_metric",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccLoggingMetric_loggingMetricLoggingBucketExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_logging_project_bucket_config" "logging_metric" {
    location  = "global"
    project   = "%{project}"
    bucket_id = "_Default"
}

resource "google_logging_metric" "logging_metric" {
  name        = "tf-test-my-(custom)/metric%{random_suffix}"
  filter      = "resource.type=gae_app AND severity>=ERROR"
  bucket_name = google_logging_project_bucket_config.logging_metric.id
}
`, context)
}

func TestAccLoggingMetric_loggingMetricDisabledExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckLoggingMetricDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccLoggingMetric_loggingMetricDisabledExample(context),
			},
			{
				ResourceName:      "google_logging_metric.logging_metric",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccLoggingMetric_loggingMetricDisabledExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_logging_metric" "logging_metric" {
  name   = "tf-test-my-(custom)/metric%{random_suffix}"
  filter = "resource.type=gae_app AND severity>=ERROR"
  metric_descriptor {
    metric_kind = "DELTA"
    value_type  = "INT64"
  }
  disabled = true
}
`, context)
}

func testAccCheckLoggingMetricDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_logging_metric" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{LoggingBasePath}}projects/{{project}}/metrics/{{%name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("LoggingMetric still exists at %s", url)
			}
		}

		return nil
	}
}
