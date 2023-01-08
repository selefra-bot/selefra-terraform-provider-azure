package resources

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
)

func GetSelefraProvider() *provider.Provider {
	diagnostics := schema.NewDiagnostics()
	selefraProvider, d := GetSelefraTerraformProvider().ToSelefraProvider(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) *bridge.TerraformBridge {
		return taskClient.(*Client).TerraformBridge
	})
	if diagnostics.AddDiagnostics(d).HasError() {
		panic(diagnostics.ToString())
	}

    selefraProvider.TableList = GetSelefraTables()

	return selefraProvider
}

func GetSelefraTables() []*schema.Table {

    diagnostics := schema.NewDiagnostics()
    tables := make([]*schema.Table, 0)
    var table *schema.Table
    var d *schema.Diagnostics

    
    table, d = TableSchemaGenerator_azurerm_maintenance_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_servicebus_namespace_network_rule_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_machine_learning_inference_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subscription_policy_remediation()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_server_extended_auditing_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_security_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_netapp_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_nginx_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_key_vault_managed_storage_account_sas_token_definition()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_nat_gateway_public_ip_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_maps_creator()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_private_link_scope()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_endpoint_storage_container()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_function_app()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_container_registry_token()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_webhook()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_hpc_cache_blob_nfs_target()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_machine_learning_workspace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_action_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_cname_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_flexible_server_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_profile()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_static_site()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_security_center_setting()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_server_security_alert_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lb_outbound_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_cluster_managed_private_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_java_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_credential()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_consumption_budget_management_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_vpn_gateway()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_security_center_assessment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_deployment_script_azure_cli()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_database_principal_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dashboard()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mariadb_virtual_network_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_hdinsight_spark_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_srv_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_managed_instance_failover_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_web_application_firewall_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_device_update_instance()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_alert_rule_machine_learning_behavior_analytics()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_output_servicebus_queue()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_express_route_circuit_peering()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_search_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_hpc_cache()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_managed_instance()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_security_center_auto_provisioning()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_azure_active_directory()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_workspace_vulnerability_assessment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_blob()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_virtual_network_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_share_dataset_kusto_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_site_recovery_fabric()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_disk_pool_iscsi_target()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channel_alexa()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_managed_application_definition()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_app()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_relay_namespace_authorization_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_source_control()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_gallery_application()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_dps_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_function_app_hybrid_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_pipeline()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_machine_learning_compute_instance()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_custom_domain()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_managed_disk_sas_token()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_parquet()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_orbital_contact_profile()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_account_customer_managed_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_attached_database_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_vpn_server_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_output_powerbi()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_windows_web_app_slot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_redis_enterprise_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_reference_input_mssql()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_share_dataset_data_lake_gen2()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_data_collection_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_resolver_dns_forwarding_ruleset()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_traffic_manager_profile()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_desktop_application()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_integration_account_partner()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_variable_bool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_site_recovery_protection_container_mapping()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_web_app_active_slot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spatial_anchors_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_windows_web_app()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_asset_filter()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_traffic_manager_external_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_public_ip()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_watchlist()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_dev_tool_portal()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_table()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_datadog_monitor_sso_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lb_backend_address_pool_address()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_public_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_signalr_shared_private_link_resource()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_signalr_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_managed_instance_vulnerability_assessment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_custom_domain_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_firewall_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_azure_table_storage()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_frontdoor_rules_engine()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_hpc_cache_blob_target()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_private_link_hub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_gateway_api()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_profile()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_active_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_application_security_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_function_javascript_uda()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_security_center_server_vulnerability_assessment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_digital_twins_endpoint_eventgrid()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_vpn_server_configuration_policy_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_disk_encryption_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_web_pubsub_hub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventgrid_domain_topic()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api_operation_tag()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_policy_virtual_machine_configuration_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_machine_scale_set_packet_capture()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channel_direct_line_speech()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_protection_backup_policy_blob_storage()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_encryption_scope()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subnet_network_security_group_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_group_user()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_recovery_services_vault()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_signalr_service_network_acl()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lighthouse_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_vmware_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logz_sub_account_tag_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_application_insights_workbook()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_packet_capture()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventhub_namespace_disaster_recovery_config()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_build_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_outbound_firewall_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iot_time_series_insights_access_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_vmware_private_cloud()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dedicated_host()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_confidential_ledger()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_environment_v3()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_advanced_threat_protection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_machine_learning_compute_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_stream_input_eventhub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_connection_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_certificate_order()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_public_ip_prefix()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_route_table()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_app_mysql_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_product()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_maintenance_assignment_virtual_machine_scale_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_protection_resource_guard()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_workflow()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_healthcare_dicom_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channel_email()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_managed_instance_transparent_data_encryption()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_orchestrated_virtual_machine_scale_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_security_center_workspace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subnet_nat_gateway_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_identity_provider_twitter()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subscription_template_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lab_service_plan()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_disk_pool_iscsi_target_lun()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_secret()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_job()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_policy_remediation()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_capacity_reservation()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_logger()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_application_insights_workbook_template()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subnet()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_streaming_locator()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_network_gateway_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_route_filter()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_container_connected_registry()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_gateway_custom_domain()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_role_definition()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_fluid_relay_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api_operation()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_shared_image_version()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_odata()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventhub_namespace_authorization_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logz_sub_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_container_registry_agent_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_route()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subscription_policy_exemption()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_runbook()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_policy_set_definition()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_link_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_database_migration_project()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventgrid_system_topic_event_subscription()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_origin_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channel_web_chat()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dns_txt_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_portal_dashboard()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_output_servicebus_topic()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_managed_instance_active_directory_administrator()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_custom_dataset()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_linux_web_app_slot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_linux_function_app()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lb_probe()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_key_vault()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_servicebus_subscription_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_slot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_machine()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventhub_namespace_customer_managed_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_firewall_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_datasource_windows_performance_counter()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_endpoint_servicebus_queue()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventhub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_identity_provider_microsoft()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cognitive_account_customer_managed_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_virtual_network_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_virtual_network_swift_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dns_ptr_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kubernetes_fleet_manager()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_firewall_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_module()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_sql_trigger()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_elasticpool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_maps_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lb_backend_address_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_hpc_cache_nfs_target()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_schedule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_iot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_gremlin_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_alert_rule_fusion()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_sql_dedicated_gateway()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_job_agent()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subscription_cost_management_export()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_backup_policy_file_share()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_diagnostic_setting()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_rule_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_policy_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_security_center_server_vulnerability_assessment_virtual_machine()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_application_insights_web_test()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_protection_backup_instance_disk()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_cassandra_datacenter()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_variable_string()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_vpn_gateway_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_gateway_certificate_authority()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_identity_provider_aad()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_trigger_recurrence()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_firewall_nat_rule_collection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_slot_custom_hostname_binding()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_hub_route_table()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_azure_file_storage()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_slot_virtual_network_swift_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_application_live_view()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_managed_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_job_credential()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_connection_monitor()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api_diagnostic()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_backup_policy_vm_workload()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_share_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iot_time_series_insights_reference_data_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dev_test_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_alert_rule_nrt()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_threat_intelligence()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_vpn_site()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_blob_inventory_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channel_sms()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_global_schema()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_desktop_scaling_plan()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_site_recovery_network_mapping()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_sql_stored_procedure()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_dynamics_365()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_origin()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_container_registry_token_password()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_attestation_provider()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_protection_backup_instance_postgresql()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_source_control_token()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_service_azure_bot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_security_center_subscription_pricing()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_action_rule_action_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_machine_scale_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_workspace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_security_center_contact()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channel_directline()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_streaming_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_gallery_application_version()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_protection_backup_policy_postgresql()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_application_gateway()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_ddos_protection_plan()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_profile()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_disk_pool_managed_disk_attachment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_named_value()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_netapp_snapshot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_job_schedule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lb_nat_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_share_dataset_kusto_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_active_directory_administrator()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_data_flow()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_consumption_budget_resource_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_hub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_management_group_policy_remediation()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api_tag_description()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_batch_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_role_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventhub_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_reference_input_blob()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_healthcare_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_flexible_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_traffic_manager_nested_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_desktop_host_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_route()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_action_rule_suppression()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subnet_service_endpoint_storage_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_output_blob()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_server_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kubernetes_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_deployment_script_azure_power_shell()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_resolver_forwarding_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_integration_account_assembly()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_datadog_monitor()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_gateway()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_job_schedule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_policy_exemption()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_machine_learning_synapse_spark()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_web_app_hybrid_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channel_facebook()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_consumption_budget_subscription()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_accelerator()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_endpoint_eventhub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_sync_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_relay_hybrid_connection_authorization_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_desktop_workspace_application_group_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_integration_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_hub_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_express_route_port()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_site_recovery_protection_container()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_enrichment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_linked_storage_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dev_test_linux_virtual_machine()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_connection_type()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_key_vault_access_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_route_disable_link_to_default_domain()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_express_route_circuit_authorization()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_watchlist_item()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_configuration_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_shared_image()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_management_group_policy_exemption()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_storage()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_managed_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_windows_virtual_machine()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_servicebus_topic_authorization_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_job()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_managed_instance_active_directory_administrator()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_active_directory_administrator()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channel_line()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_alert_processing_rule_suppression()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_route_map()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_microsoft_cloud_app_security()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_container()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_management_group_template_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_trigger_custom()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_databox_edge_order()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_servicebus_queue()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventhub_authorization_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_alert_rule_scheduled()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_txt_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_hub_ip()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_batch_job()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mariadb_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_scheduled_query_rules_log()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_container_registry_scope_map()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logz_tag_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dedicated_host_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_healthcare_medtech_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_healthbot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_zone()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_flexible_server_firewall_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_aad_diagnostic_setting()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_workspace_aad_admin()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_netapp_snapshot_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_frontdoor_firewall_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_local_network_gateway()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_a_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_integration_runtime_managed()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_flexible_server_firewall_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_consumer_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subscription_policy_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_server_transparent_data_encryption()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dev_test_windows_virtual_machine()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_vmware_netapp_volume_attachment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_managed_private_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_office_365()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_disk_access()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_odbc()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dns_cname_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_databox_edge_device()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_aws_cloud_trail()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_notification_hub_authorization_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_hdinsight_interactive_query_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_ip_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_data_collection_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_backup_protected_file_share()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_cluster_principal_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_custom_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_transform()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_integration_account_batch_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_video_analyzer()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_orbital_spacecraft()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_servicebus_topic()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_saved_search()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_shared_access_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_powerbi_embedded()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_integration_runtime_self_hosted()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_trigger_custom_event()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_microsoft_threat_protection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_flexible_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_cassandra_table()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api_operation_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_sql_pool_vulnerability_assessment_baseline()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_office_365_project()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_windows_function_app()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_interface_application_security_group_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_integration_account_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dev_test_virtual_network()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_snowflake()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_spark_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventgrid_topic()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_management_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_output_mssql()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_key_vault_certificate_issuer()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_proximity_placement_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_source_control()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channels_registration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_postgresql()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_netapp_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_ptr_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_video_analyzer_edge_module()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_container_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_build_pack_binding()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dns_aaaa_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dev_test_lab()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dns_ns_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_sql_role_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_custom_domain()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_cosmosdb_sqlapi()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_provider_registration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_backend()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_network()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iot_time_series_insights_standard_environment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channel_ms_teams()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_managed_application()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_vmware_express_route_authorization()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_healthcare_fhir_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_availability_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_express_route_circuit_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_stream_input_iothub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_static_site_custom_domain()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_product_api()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_product_tag()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_query_pack_query()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_protection_backup_vault()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_data_collection_rule_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_databricks_workspace_customer_managed_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_managed_instance_security_alert_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_mx_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dedicated_hardware_security_module()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_flexible_server_active_directory_administrator()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_service_fabric_managed_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_flexible_server_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dns_srv_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_security_center_assessment_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_trigger_http_request()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_elastic_cloud_elasticsearch()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_identity_provider_aadb2c()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_resolver()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_managed_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cognitive_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_hub_route_table_route()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_trigger_blob_event()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_account_network_rules()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_custom_domain()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_cluster_customer_managed_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_interface_nat_rule_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_security_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_app_redis_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_alert_rule_ms_security_incident()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_private_link_scoped_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_frontdoor()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_sql_pool_extended_auditing_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_sql_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_source_control_token()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_sql_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_mysql()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_firewall_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_firewall_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_stream_input_eventhub_v2()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_nat_gateway()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_group_policy_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dev_test_global_vm_shutdown_schedule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_ssh_public_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_hub_security_partner_provider()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_azure_blob_storage()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_group_cost_management_export()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_mongo_collection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_nat_gateway_public_ip_prefix_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_managed_private_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_aws_s3()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_action_http()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_web_pubsub_shared_private_link_resource()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_dsc_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_digital_twins_instance()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_environment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_linux_virtual_machine()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lighthouse_definition()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_datadog_monitor_tag_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_site_recovery_replicated_vm()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_database_vulnerability_assessment_rule_baseline()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_certificate_binding()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_cluster_customer_managed_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_key_vault_managed_storage_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_resolver_outbound_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_express_route_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_firewall_network_rule_collection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api_release()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_application_insights()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_dsc_nodeconfiguration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_datasource_windows_event()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_log_profile()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_network_peering()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_data_lake_gen2_path()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_workspace_sql_aad_admin()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_sql_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_firewall_application_rule_collection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_cassandra_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_container_registry_task()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_app_cosmosdb_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_application_insights_smart_detection_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dashboard_grafana()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_share_file()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_activity_log_alert()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_standard()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_sql_pool_workload_classifier()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_linked_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_capacity_reservation_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_watcher()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mariadb_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_query_pack()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_server_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_software_update_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_aaaa_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_office_power_bi()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_notification_hub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_queue()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_workspace_security_alert_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventgrid_event_subscription()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_synapse()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_product_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_source_control_slot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_virtual_network_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_object_replication()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_databricks_access_connector()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_management_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_databricks_workspace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_web_pubsub_network_acl()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_express_route_circuit()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_solution()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_script()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_notification_recipient_email()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dev_test_schedule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_output_function()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_web_app()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_mongo_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_express_route_gateway()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_share()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_variable_int()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_sql_container()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_server_dns_alias()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_application_insights_api_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_backup_protected_vm()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_interface_backend_address_pool_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_batch_application()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_portal_tenant_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_shared_image_gallery()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_share_directory()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_redis_cache()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_sql_role_definition()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_output_synapse()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_sync()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_healthcare_workspace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_key_vault_secret()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_office_irm()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_stream_input_blob()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_integration_account_session()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_api_portal_custom_domain()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iot_time_series_insights_gen2_environment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_batch_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kubernetes_cluster_node_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_network_dns_servers()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_http()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_failover_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_managed_private_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_manager()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_dps()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_output_eventhub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_route_server_bgp_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_web_pubsub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_web()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_watcher()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_integration_runtime_self_hosted()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_virtual_machine()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_custom_hostname_binding()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_binary()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_nginx_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_integration_runtime_azure_ssis()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_active_directory_domain_service_trust()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_fallback_route()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subscription()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_backup_policy_vm()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_active_slot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_management_lock()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_metric_alert()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_machine_extension()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_virtual_network_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dns_caa_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_snowflake()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_redis_cache()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_sql_pool_security_alert_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_openid_connect_provider()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventhub_namespace_schema_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_cosmosdb()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_share_dataset_blob_storage()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_content_key_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_iothub_data_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_billing_account_cost_management_export()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_relay_hybrid_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_firewall()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stack_hci_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_workspace_extended_auditing_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_wan()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_trigger_tumbling_window()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventhub_namespace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_blueprint_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_communication_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_zone_virtual_network_link()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_image()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventgrid_system_topic()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_network_gateway()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_aadb2c_directory()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_healthcare_medtech_service_fhir_destination()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_tag()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_subscription()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_postgresql()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_gateway()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_data_lake_storage_gen2()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lb_nat_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_marketplace_agreement()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_servicebus_subscription()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_maintenance_assignment_virtual_machine()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_custom_provider()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_delimited_text()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_sftp()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_failover_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_cassandra_keyspace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_notebook_workspace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_output_table()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_action_custom()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_hdinsight_hbase_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lb_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_linux_web_app()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_function_app_slot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_lb()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_key_vault_managed_hardware_security_module()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iotcentral_application_network_rule_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_builder()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_hdinsight_kafka_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_load_test()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_analysis_services_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_protection_backup_instance_blob_storage()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_service_fabric_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_json()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_key_vault_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_firewall_policy_rule_collection_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mariadb_firewall_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_container_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_active_directory_administrator()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_elasticpool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_relay_namespace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_scheduled_query_rules_alert()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_managed_instance()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_sql_pool_workload_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_key_vault()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_protection_backup_policy_disk()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dns_zone()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_linux_function_app_slot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_smart_detector_alert_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_route_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_firewall_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_gateway_route_config()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_resolver_virtual_network_link()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_endpoint_custom_domain()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_interface_application_gateway_backend_address_pool_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_scheduled_query_rules_alert_v2()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_variable_datetime()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_autoscale_setting()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mysql_flexible_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_email_template()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_digital_twins_endpoint_servicebus()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_snapshot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_hybrid_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api_version_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_watcher_flow_log()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_machine_packet_capture()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_machine_scale_set_extension()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_tenant_template_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_gateway_host_name_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_sql_server_table()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_mysql()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_live_event()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_servicebus_namespace_disaster_recovery_config()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_key_vault_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_hybrid_runbook_worker_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_network_gateway_nat_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_sql_function()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_frontdoor_custom_https_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_identity_provider_google()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_automation_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_function_javascript_udf()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bot_channel_slack()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventgrid_domain()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_windows_function_app_slot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_traffic_manager_azure_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_authorization_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_servicebus_queue_authorization_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_server_vulnerability_assessment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cdn_frontdoor_security_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_microsoft_defender_advanced_threat_protection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_identity_provider_facebook()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_configuration_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_flowlet_data_flow()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_container_registry_task_schedule_run_now()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_dps_shared_access_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api_tag()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_managed_disk()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_database_extended_auditing_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_table()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_function_app_active_slot()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_machine_data_disk_attachment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_template_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_group_template_deployment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_azure_databricks()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iot_time_series_insights_event_source_iothub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_user()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_data_lake_gen2_filesystem()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_workspace_key()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_container_registry_webhook()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_stream_analytics_output_cosmosdb()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_storage_insights()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_windows_virtual_machine_scale_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_desktop_host_pool_registration_info()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_server_microsoft_support_auditing_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_interface_security_group_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_maintenance_assignment_dedicated_host()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_cosmosdb_mongoapi()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_services_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_integration_account_map()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_service_plan()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_hpc_cache_access_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_servicebus_namespace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_connection_service_principal()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_management_group_subscription_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cognitive_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_manager_network_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_desktop_application_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_api_schema()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_site_recovery_replication_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_disk_pool()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_private_dns_resolver_inbound_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_redis_enterprise_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iot_security_solution()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_active_directory_domain_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_integration_service_environment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iotcentral_application()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_monitor_alert_processing_rule_action_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_function_app_function()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_digital_twins_endpoint_eventhub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_redis_linked_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_redis_firewall_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_configuration_feature()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_connection_classic_certificate()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_federated_identity_credential()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_postgresql_flexible_server_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_group_policy_exemption()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_nginx_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_netapp_volume()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sql_server()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mariadb_configuration()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_azure_advanced_threat_protection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_hdinsight_hadoop_cluster()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_vpn_gateway_nat_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_customized_accelerator()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_role_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_automation_hybrid_runbook_worker()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_hub_bgp_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_subnet_route_table_association()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_eventhub_consumer_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iot_time_series_insights_event_source_eventhub()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_streaming_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iot_security_device_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_workspace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_office_atp()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_integration_runtime_azure()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_linked_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_eventhub_data_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_resource_group_policy_remediation()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_digital_twins_time_series_database_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_spring_cloud_api_portal()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dns_a_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_bastion_host()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_point_to_site_vpn_gateway()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_app_service_plan()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_dataset_azure_blob()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_application_insights_analytics_item()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_diagnostic()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_azure_sql_database()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_sentinel_data_connector_azure_security_center()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_purview_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_security_center_automation()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_azure_function()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_network_interface()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_endpoint_servicebus_topic()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logz_monitor()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_kusto_eventgrid_data_connection()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_azure_search()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_linked_service_kusto()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_dns_mx_record()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_notification_recipient_user()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_batch_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_integration_account_agreement()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_key_vault_certificate_contacts()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_asset()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_sql_pool_vulnerability_assessment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_linux_virtual_machine_scale_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_factory_trigger_schedule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_sync_cloud_endpoint()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_backup_container_storage_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_data_share()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_synapse_integration_runtime_azure()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_route()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_iothub_device_update_account()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_management_group_policy_assignment()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_virtual_desktop_workspace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_database_migration_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_mssql_managed_instance_failover_group()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_policy_definition()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_active_directory_domain_service_replica_set()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_servicebus_namespace_authorization_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_firewall_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_log_analytics_data_export_rule()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_notification_hub_namespace()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_search_shared_private_link_service()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_container_registry()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_media_live_event_output()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_logic_app_integration_account_schema()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_user_assigned_identity()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_cosmosdb_gremlin_graph()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_api_management_product_policy()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    
    table, d = TableSchemaGenerator_azurerm_storage_table_entity()
    if !diagnostics.AddDiagnostics(d).HasError() {
        tables = append(tables, table)
    }
    

    if diagnostics.HasError() {
        panic(diagnostics.ToString())
    }

	return tables
}
