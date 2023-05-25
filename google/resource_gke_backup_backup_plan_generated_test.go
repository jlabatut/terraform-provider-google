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

func TestAccGKEBackupBackupPlan_gkebackupBackupplanBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEBackupBackupPlanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupBackupPlan_gkebackupBackupplanBasicExample(context),
			},
			{
				ResourceName:            "google_gke_backup_backup_plan.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccGKEBackupBackupPlan_gkebackupBackupplanBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-basic-cluster%{random_suffix}"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-basic-plan%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}
`, context)
}

func TestAccGKEBackupBackupPlan_gkebackupBackupplanAutopilotExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEBackupBackupPlanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupBackupPlan_gkebackupBackupplanAutopilotExample(context),
			},
			{
				ResourceName:            "google_gke_backup_backup_plan.autopilot",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccGKEBackupBackupPlan_gkebackupBackupplanAutopilotExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-autopilot-cluster%{random_suffix}"
  location           = "us-central1"
  enable_autopilot = true
  ip_allocation_policy {   
  }
  release_channel {
    channel = "RAPID"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
}

resource "google_gke_backup_backup_plan" "autopilot" {
  name = "tf-test-autopilot-plan%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}
`, context)
}

func TestAccGKEBackupBackupPlan_gkebackupBackupplanCmekExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEBackupBackupPlanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupBackupPlan_gkebackupBackupplanCmekExample(context),
			},
			{
				ResourceName:            "google_gke_backup_backup_plan.cmek",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccGKEBackupBackupPlan_gkebackupBackupplanCmekExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-cmek-cluster%{random_suffix}"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
}

resource "google_gke_backup_backup_plan" "cmek" {
  name = "tf-test-cmek-plan%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    selected_namespaces {
      namespaces = ["default", "test"]
    }
    encryption_key {
      gcp_kms_encryption_key = google_kms_crypto_key.crypto_key.id
    }
  }
}

resource "google_kms_crypto_key" "crypto_key" {
  name     = "tf-test-backup-key%{random_suffix}"
  key_ring = google_kms_key_ring.key_ring.id
}

resource "google_kms_key_ring" "key_ring" {
  name     = "tf-test-backup-key%{random_suffix}"
  location = "us-central1"
}
`, context)
}

func TestAccGKEBackupBackupPlan_gkebackupBackupplanFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project":       acctest.GetTestProjectFromEnv(),
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckGKEBackupBackupPlanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupBackupPlan_gkebackupBackupplanFullExample(context),
			},
			{
				ResourceName:            "google_gke_backup_backup_plan.full",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccGKEBackupBackupPlan_gkebackupBackupplanFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-full-cluster%{random_suffix}"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
}

resource "google_gke_backup_backup_plan" "full" {
  name = "tf-test-full-plan%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  retention_policy {
    backup_delete_lock_days = 30
    backup_retain_days = 180
  }
  backup_schedule {
    cron_schedule = "0 9 * * 1"
  }
  backup_config {
    include_volume_data = true
    include_secrets = true
    selected_applications {
      namespaced_names {
        name = "app1"
        namespace = "ns1"
      }
      namespaced_names {
        name = "app2"
        namespace = "ns2"
      }
    }
  }
}
`, context)
}

func testAccCheckGKEBackupBackupPlanDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_gke_backup_backup_plan" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{GKEBackupBasePath}}projects/{{project}}/locations/{{location}}/backupPlans/{{name}}")
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
				return fmt.Errorf("GKEBackupBackupPlan still exists at %s", url)
			}
		}

		return nil
	}
}
