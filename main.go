package main
// #include <stdint.h>     
// #include <stdbool.h>    
// #include <stdio.h>      
// #include <string.h>     
// #include <signal.h>     
// #include <time.h>       
// #include <unistd.h>     
// #include <stdlib.h>     
// #include "parson.h"
// #include "loragw_hal.h"
// static void sig_handler(int sigio) {
//     if (sigio == SIGQUIT) {
//         quit_sig = 1;;
//     } else if ((sigio == SIGINT) || (sigio == SIGTERM)) {
//         exit_sig = 1;
//     }
// }

// int parse_SX1301_configuration(const char * conf_file) {
//     int i;
//     const char conf_obj[] = "SX1301_conf";
//     char param_name[32]; 
//     const char *str; 
//     struct lgw_conf_board_s boardconf;
//     struct lgw_conf_rxrf_s rfconf;
//     struct lgw_conf_rxif_s ifconf;
//     JSON_Value *root_val;
//     JSON_Object *root = NULL;
//     JSON_Object *conf = NULL;
//     JSON_Value *val;
//     uint32_t sf, bw;

//    
//     root_val = json_parse_file_with_comments(conf_file);
//     root = json_value_get_object(root_val);
//     if (root == NULL) {
//         MSG("ERROR: %s id not a valid JSON file\n", conf_file);
//         exit(EXIT_FAILURE);
//     }
//     conf = json_object_get_object(root, conf_obj);
//     if (conf == NULL) {
//         MSG("INFO: %s does not contain a JSON object named %s\n", conf_file, conf_obj);
//         return -1;
//     } else {
//         MSG("INFO: %s does contain a JSON object named %s, parsing SX1301 parameters\n", conf_file, conf_obj);
//     }
//     memset(&boardconf, 0, sizeof boardconf); 
//     val = json_object_get_value(conf, "lorawan_public");
//     if (json_value_get_type(val) == JSONBoolean) {
//         boardconf.lorawan_public = (bool)json_value_get_boolean(val);
//     } else {
//         MSG("WARNING: Data type for lorawan_public seems wrong, please check\n");
//         boardconf.lorawan_public = false;
//     }
//     val = json_object_get_value(conf, "clksrc"); 
//     if (json_value_get_type(val) == JSONNumber) {
//         boardconf.clksrc = (uint8_t)json_value_get_number(val);
//     } else {
//         MSG("WARNING: Data type for clksrc seems wrong, please check\n");
//         boardconf.clksrc = 0;
//     }
//     MSG("INFO: lorawan_public %d, clksrc %d\n", boardconf.lorawan_public, boardconf.clksrc);
//     if (lgw_board_setconf(boardconf) != LGW_HAL_SUCCESS) {
//         MSG("ERROR: Failed to configure board\n");
//         return -1;
//     }

//     for (i = 0; i < LGW_RF_CHAIN_NB; ++i) {
//         memset(&rfconf, 0, sizeof(rfconf)); 
//         sprintf(param_name, "radio_%i", i); 
//         val = json_object_get_value(conf, param_name); 
//         if (json_value_get_type(val) != JSONObject) {
//             MSG("INFO: no configuration for radio %i\n", i);
//             continue;
//         }
// 
//         sprintf(param_name, "radio_%i.enable", i);
//         val = json_object_dotget_value(conf, param_name);
//         if (json_value_get_type(val) == JSONBoolean) {
//             rfconf.enable = (bool)json_value_get_boolean(val);
//         } else {
//             rfconf.enable = false;
//         }
//         if (rfconf.enable == false) { 
//             MSG("INFO: radio %i disabled\n", i);
//         } else  { 
//             snprintf(param_name, sizeof param_name, "radio_%i.freq", i);
//             rfconf.freq_hz = (uint32_t)json_object_dotget_number(conf, param_name);
//             snprintf(param_name, sizeof param_name, "radio_%i.rssi_offset", i);
//             rfconf.rssi_offset = (float)json_object_dotget_number(conf, param_name);
//             snprintf(param_name, sizeof param_name, "radio_%i.type", i);
//             str = json_object_dotget_string(conf, param_name);
//             if (!strncmp(str, "SX1255", 6)) {
//                 rfconf.type = LGW_RADIO_TYPE_SX1255;
//             } else if (!strncmp(str, "SX1257", 6)) {
//                 rfconf.type = LGW_RADIO_TYPE_SX1257;
//             } else {
//                 MSG("WARNING: invalid radio type: %s (should be SX1255 or SX1257)\n", str);
//             }
//             snprintf(param_name, sizeof param_name, "radio_%i.tx_enable", i);
//             val = json_object_dotget_value(conf, param_name);
//             if (json_value_get_type(val) == JSONBoolean) {
//                 rfconf.tx_enable = (bool)json_value_get_boolean(val);
//                 if (rfconf.tx_enable == true) {
//                     
//                     snprintf(param_name, sizeof param_name, "radio_%i.tx_notch_freq", i);
//                     rfconf.tx_notch_freq = (uint32_t)json_object_dotget_number(conf, param_name);
//                 }
//             } else {
//                 rfconf.tx_enable = false;
//             }
//             MSG("INFO: radio %i enabled (type %s), center frequency %u, RSSI offset %f, tx enabled %d, tx_notch_freq %u\n", i, str, rfconf.freq_hz, rfconf.rssi_offset, rfconf.tx_enable, rfconf.tx_notch_freq);
//         }
//         if (lgw_rxrf_setconf(i, rfconf) != LGW_HAL_SUCCESS) {
//             MSG("ERROR: invalid configuration for radio %i\n", i);
//             return -1;
//         }
//     }

//     
//     for (i = 0; i < LGW_MULTI_NB; ++i) {
//         memset(&ifconf, 0, sizeof(ifconf)); 
//         sprintf(param_name, "chan_multiSF_%i", i); 
//         val = json_object_get_value(conf, param_name); 
//         if (json_value_get_type(val) != JSONObject) {
//             MSG("INFO: no configuration for LoRa multi-SF channel %i\n", i);
//             continue;
//         }
//         sprintf(param_name, "chan_multiSF_%i.enable", i);
//         val = json_object_dotget_value(conf, param_name);
//         if (json_value_get_type(val) == JSONBoolean) {
//             ifconf.enable = (bool)json_value_get_boolean(val);
//         } else {
//             ifconf.enable = false;
//         }
//         if (ifconf.enable == false) { 
//             MSG("INFO: LoRa multi-SF channel %i disabled\n", i);
//         } else  {
//             sprintf(param_name, "chan_multiSF_%i.radio", i);
//             ifconf.rf_chain = (uint32_t)json_object_dotget_number(conf, param_name);
//             sprintf(param_name, "chan_multiSF_%i.if", i);
//             ifconf.freq_hz = (int32_t)json_object_dotget_number(conf, param_name);
//             // TODO: handle individual SF enabling and disabling (spread_factor)
//             MSG("INFO: LoRa multi-SF channel %i enabled, radio %i selected, IF %i Hz, 125 kHz bandwidth, SF 7 to 12\n", i, ifconf.rf_chain, ifconf.freq_hz);
//         }
//         
//         if (lgw_rxif_setconf(i, ifconf) != LGW_HAL_SUCCESS) {
//             MSG("ERROR: invalid configuration for Lora multi-SF channel %i\n", i);
//             return -1;
//         }
//     }

//     memset(&ifconf, 0, sizeof(ifconf)); 
//     val = json_object_get_value(conf, "chan_Lora_std"); 
//     if (json_value_get_type(val) != JSONObject) {
//         MSG("INFO: no configuration for LoRa standard channel\n");
//     } else {
//         val = json_object_dotget_value(conf, "chan_Lora_std.enable");
//         if (json_value_get_type(val) == JSONBoolean) {
//             ifconf.enable = (bool)json_value_get_boolean(val);
//         } else {
//             ifconf.enable = false;
//         }
//         if (ifconf.enable == false) {
//             MSG("INFO: LoRa standard channel %i disabled\n", i);
//         } else  {
//             ifconf.rf_chain = (uint32_t)json_object_dotget_number(conf, "chan_Lora_std.radio");
//             ifconf.freq_hz = (int32_t)json_object_dotget_number(conf, "chan_Lora_std.if");
//             bw = (uint32_t)json_object_dotget_number(conf, "chan_Lora_std.bandwidth");
//             switch(bw) {
//                 case 500000: ifconf.bandwidth = BW_500KHZ; break;
//                 case 250000: ifconf.bandwidth = BW_250KHZ; break;
//                 case 125000: ifconf.bandwidth = BW_125KHZ; break;
//                 default: ifconf.bandwidth = BW_UNDEFINED;
//             }
//             sf = (uint32_t)json_object_dotget_number(conf, "chan_Lora_std.spread_factor");
//             switch(sf) {
//                 case  7: ifconf.datarate = DR_LORA_SF7;  break;
//                 case  8: ifconf.datarate = DR_LORA_SF8;  break;
//                 case  9: ifconf.datarate = DR_LORA_SF9;  break;
//                 case 10: ifconf.datarate = DR_LORA_SF10; break;
//                 case 11: ifconf.datarate = DR_LORA_SF11; break;
//                 case 12: ifconf.datarate = DR_LORA_SF12; break;
//                 default: ifconf.datarate = DR_UNDEFINED;
//             }
//             MSG("INFO: LoRa standard channel enabled, radio %i selected, IF %i Hz, %u Hz bandwidth, SF %u\n", ifconf.rf_chain, ifconf.freq_hz, bw, sf);
//         }
//         if (lgw_rxif_setconf(8, ifconf) != LGW_HAL_SUCCESS) {
//             MSG("ERROR: invalid configuration for Lora standard channel\n");
//             return -1;
//         }
//     }

// 
//     memset(&ifconf, 0, sizeof(ifconf)); 
//     val = json_object_get_value(conf, "chan_FSK"); 
//     if (json_value_get_type(val) != JSONObject) {
//         MSG("INFO: no configuration for FSK channel\n");
//     } else {
//         val = json_object_dotget_value(conf, "chan_FSK.enable");
//         if (json_value_get_type(val) == JSONBoolean) {
//             ifconf.enable = (bool)json_value_get_boolean(val);
//         } else {
//             ifconf.enable = false;
//         }
//         if (ifconf.enable == false) {
//             MSG("INFO: FSK channel %i disabled\n", i);
//         } else  {
//             ifconf.rf_chain = (uint32_t)json_object_dotget_number(conf, "chan_FSK.radio");
//             ifconf.freq_hz = (int32_t)json_object_dotget_number(conf, "chan_FSK.if");
//             bw = (uint32_t)json_object_dotget_number(conf, "chan_FSK.bandwidth");
//             if      (bw <= 7800)   ifconf.bandwidth = BW_7K8HZ;
//             else if (bw <= 15600)  ifconf.bandwidth = BW_15K6HZ;
//             else if (bw <= 31200)  ifconf.bandwidth = BW_31K2HZ;
//             else if (bw <= 62500)  ifconf.bandwidth = BW_62K5HZ;
//             else if (bw <= 125000) ifconf.bandwidth = BW_125KHZ;
//             else if (bw <= 250000) ifconf.bandwidth = BW_250KHZ;
//             else if (bw <= 500000) ifconf.bandwidth = BW_500KHZ;
//             else ifconf.bandwidth = BW_UNDEFINED;
//             ifconf.datarate = (uint32_t)json_object_dotget_number(conf, "chan_FSK.datarate");
//             MSG("INFO: FSK channel enabled, radio %i selected, IF %i Hz, %u Hz bandwidth, %u bps datarate\n", ifconf.rf_chain, ifconf.freq_hz, bw, ifconf.datarate);
//         }
//         if (lgw_rxif_setconf(9, ifconf) != LGW_HAL_SUCCESS) {
//             MSG("ERROR: invalid configuration for FSK channel\n");
//             return -1;
//         }
//     }
//     json_value_free(root_val);
//     return 0;
// }

// int parse_gateway_configuration(const char * conf_file) {
//     const char conf_obj[] = "gateway_conf";
//     JSON_Value *root_val;
//     JSON_Object *root = NULL;
//     JSON_Object *conf = NULL;
//     const char *str; 

//     
//     root_val = json_parse_file_with_comments(conf_file);
//     root = json_value_get_object(root_val);
//     if (root == NULL) {
//         MSG("ERROR: %s id not a valid JSON file\n", conf_file);
//         exit(EXIT_FAILURE);
//     }
//     conf = json_object_get_object(root, conf_obj);
//     if (conf == NULL) {
//         MSG("INFO: %s does not contain a JSON object named %s\n", conf_file, conf_obj);
//         return -1;
//     } else {
//         MSG("INFO: %s does contain a JSON object named %s, parsing gateway parameters\n", conf_file, conf_obj);
//     }

//     str = json_object_get_string(conf, "gateway_ID");
//     if (str != NULL) {
//         sscanf(str, "%llx", &ull);
//         lgwm = ull;
//         MSG("INFO: gateway MAC address is configured to %016llX\n", ull);
//     }

//     json_value_free(root_val);
//     return 0;
// }

// void open_log(void) {
//     int i;
//     char iso_date[20];

//     strftime(iso_date,ARRAY_SIZE(iso_date),"%Y%m%dT%H%M%SZ",gmtime(&now_time)); 

//     sprintf(log_file_name, "pktlog_%s_%s.csv", lgwm_str, iso_date);
//     log_file = fopen(log_file_name, "a"); 
//     if (log_file == NULL) {
//         MSG("ERROR: impossible to create log file %s\n", log_file_name);
//         exit(EXIT_FAILURE);
//     }

//     i = fprintf(log_file, "\"gateway ID\",\"node MAC\",\"UTC timestamp\",\"us count\",\"frequency\",\"RF chain\",\"RX chain\",\"status\",\"size\",\"modulation\",\"bandwidth\",\"datarate\",\"coderate\",\"RSSI\",\"SNR\",\"payload\"\n");
//     if (i < 0) {
//         MSG("ERROR: impossible to write to log file %s\n", log_file_name);
//         exit(EXIT_FAILURE);
//     }

//     MSG("INFO: Now writing to log file %s\n", log_file_name);
//     return;
// }

// void usage(void) {
//	printf("*** Library version information ***\n%s\n\n", lgw_version_info());
//     printf( "Available options:\n");
//     printf( " -h print this help\n");
//     printf( " -r <int> rotate log file every N seconds (-1 disable log rotation)\n");
// }
import "C"
func main(){
	var i, j int
	var timespec sleep_time = {0, 3000000}
	
	var log_rotate_interval int = 3600
	var  time_check int = 0
	var pkt_in_log uint64 = 0
	const global_conf_fname[] char = "global_conf.json"
    const local_conf_fname[] char = "local_conf.json"
	const debug_conf_fname[] char = "debug_conf.json"
	
	var rxpkt[16] lgw_pkt_rx_s
	var p *lgw_pkt_rx_s

	var nb_pkt int
	var fetch_time timespec
	var fetch_timestamp[30]
	var x *tm
	for (i = getopt (argc, argv, "hr:")) != -1{
		switch i{
		case 'h':
			usage()
			return EXIT_FAILURE
			break
		case 'r':
			log_rotate_interval = atoi(optarg)
			if (log_rotate_interval == 0) || (log_rotate_interval < -1) {
				MSG( "ERROR: Invalid argument for -r option\n")
				return EXIT_FAILURE
			}
			break;

		default:
			MSG("ERROR: argument parsing use -h option for help\n");
			usage();
			return EXIT_FAILURE;
		}
	}

	sigemptyset(&sigact.sa_mask);
    sigact.sa_flags = 0;
    sigact.sa_handler = sig_handler;
    sigaction(SIGQUIT, &sigact, NULL);
    sigaction(SIGINT, &sigact, NULL);
	sigaction(SIGTERM, &sigact, NULL);
	
	if access(debug_conf_fname, R_OK) == 0{
		MSG("INFO: found debug configuration file %s, other configuration files will be ignored\n", debug_conf_fname);
        parse_SX1301_configuration(debug_conf_fname);
        parse_gateway_configuration(debug_conf_fname);
	}else if access(global_conf_fname, R_OK) == 0{
		MSG("INFO: found global configuration file %s, trying to parse it\n", global_conf_fname);
        parse_SX1301_configuration(global_conf_fname);
		parse_gateway_configuration(global_conf_fname);
		if access(local_conf_fname, R_OK) == 0{
            MSG("INFO: found local configuration file %s, trying to parse it\n", local_conf_fname);
            parse_SX1301_configuration(local_conf_fname);
            parse_gateway_configuration(local_conf_fname);
		}
	}else if access(local_conf_fname, R_OK) == 0 {
		MSG("INFO: found local configuration file %s, trying to parse it\n", local_conf_fname);
        parse_SX1301_configuration(local_conf_fname);
        parse_gateway_configuration(local_conf_fname);
	}else {
        MSG("ERROR: failed to find any configuration file named %s, %s or %s\n", global_conf_fname, local_conf_fname, debug_conf_fname);
		return EXIT_FAILURE;
	}

	i =lgw_start();
	if i == LGW_HAL_SUCCESS {
        MSG("INFO: concentrator started, packet can now be received\n");
    } else {
        MSG("ERROR: failed to start the concentrator\n");
        return EXIT_FAILURE;
	}
	
	sprintf(lgwm_str, "%08X%08X", (uint32_t)(lgwm >> 32), (uint32_t)(lgwm & 0xFFFFFFFF));

	time(&now_time);
	open_log();
	
	for (quit_sig != 1) && (exit_sig != 1){
		nb_pkt = lgw_receive(ARRAY_SIZE(rxpkt), rxpkt);
        if nb_pkt == LGW_HAL_ERROR {
            MSG("ERROR: failed packet fetch, exiting\n");
            return EXIT_FAILURE;
        } else if nb_pkt == 0 {
            clock_nanosleep(CLOCK_MONOTONIC, 0, &sleep_time, NULL); 
        } else {
           
            clock_gettime(CLOCK_REALTIME, &fetch_time);
            x = gmtime(&(fetch_time.tv_sec));
            sprintf(fetch_timestamp,"%04i-%02i-%02i %02i:%02i:%02i.%03liZ",(x->tm_year)+1900,(x->tm_mon)+1,x->tm_mday,x->tm_hour,x->tm_min,x->tm_sec,(fetch_time.tv_nsec)/1000000); 
        }

	for i=0; i < nb_pkt; ++i {
		p ：= &rxpkt[i];

		fprintf(log_file, "\"%08X%08X\",", (uint32_t)(lgwm >> 32), (uint32_t)(lgwm & 0xFFFFFFFF));
		fputs("\"\",", log_file); 
		fprintf(log_file, "\"%s\",", fetch_timestamp);
		fprintf(log_file, "%10u,", p->count_us);
		fprintf(log_file, "%10u,", p->freq_hz);
		fprintf(log_file, "%u,", p->rf_chain);
		fprintf(log_file, "%2d,", p->if_chain);
		switch p->status{
			case STAT_CRC_OK:       fputs("\"CRC_OK\" ,", log_file); break;
			case STAT_CRC_BAD:      fputs("\"CRC_BAD\",", log_file); break;
			case STAT_NO_CRC:       fputs("\"NO_CRC\" ,", log_file); break;
			case STAT_UNDEFINED:    fputs("\"UNDEF\"  ,", log_file); break;
			default:                fputs("\"ERR\"    ,", log_file);
		}
		fprintf(log_file, "%3u,", p->size);
		switch p->modulation {
			case MOD_LORA:  fputs("\"LORA\",", log_file); break;
			case MOD_FSK:   fputs("\"FSK\" ,", log_file); break;
			default:        fputs("\"ERR\" ,", log_file);
		}
		switch p->bandwidth {
			case BW_500KHZ:     fputs("500000,", log_file); break;
			case BW_250KHZ:     fputs("250000,", log_file); break;
			case BW_125KHZ:     fputs("125000,", log_file); break;
			case BW_62K5HZ:     fputs("62500 ,", log_file); break;
			case BW_31K2HZ:     fputs("31200 ,", log_file); break;
			case BW_15K6HZ:     fputs("15600 ,", log_file); break;
			case BW_7K8HZ:      fputs("7800  ,", log_file); break;
			case BW_UNDEFINED:  fputs("0     ,", log_file); break;
			default:            fputs("-1    ,", log_file);
		}

		if p->modulation == MOD_LORA {
			switch p->datarate {
				case DR_LORA_SF7:   fputs("\"SF7\"   ,", log_file); break;
				case DR_LORA_SF8:   fputs("\"SF8\"   ,", log_file); break;
				case DR_LORA_SF9:   fputs("\"SF9\"   ,", log_file); break;
				case DR_LORA_SF10:  fputs("\"SF10\"  ,", log_file); break;
				case DR_LORA_SF11:  fputs("\"SF11\"  ,", log_file); break;
				case DR_LORA_SF12:  fputs("\"SF12\"  ,", log_file); break;
				default:            fputs("\"ERR\"   ,", log_file);
			}
		} else if p->modulation == MOD_FSK {
			fprintf(log_file, "\"%6u\",", p->datarate);
		} else {
			fputs("\"ERR\"   ,", log_file);
		}

		switch p->coderate {
			case CR_LORA_4_5:   fputs("\"4/5\",", log_file); break;
			case CR_LORA_4_6:   fputs("\"2/3\",", log_file); break;
			case CR_LORA_4_7:   fputs("\"4/7\",", log_file); break;
			case CR_LORA_4_8:   fputs("\"1/2\",", log_file); break;
			case CR_UNDEFINED:  fputs("\"\"   ,", log_file); break;
			default:            fputs("\"ERR\",", log_file);
		}

		fprintf(log_file, "%+.0f,", p->rssi);

		fprintf(log_file, "%+5.1f,", p->snr);

		fputs("\"", log_file);
		for j = 0; j < p->size; ++j {
			if (j > 0) && (j%4 == 0)){
				fputs("-", log_file);
			} 
			fprintf(log_file, "%02X", p->payload[j]);
		}

		fputs("\"\n", log_file);
		fflush(log_file);
		pkt_in_log=1+pkt_in_log;
	}

	time_check=1+time_check

	if time_check >= 8 {
		time_check = 0;
		time(&now_time);
		if difftime(now_time, log_start_time) > log_rotate_interval {
			fclose(log_file);
			MSG("INFO: log file %s closed, %lu packet(s) recorded\n", log_file_name, pkt_in_log);
			pkt_in_log = 0;
			open_log();
		}
	}
}
	if exit_sig == 1 {
        
        i = lgw_stop();
        if i == LGW_HAL_SUCCESS {
            MSG("INFO: concentrator stopped successfully\n");
        } else {
            MSG("WARNING: failed to stop concentrator successfully\n");
        }
        fclose(log_file);
        MSG("INFO: log file %s closed, %lu packet(s) recorded\n", log_file_name, pkt_in_log);
	}
	
	MSG("INFO: Exiting packet logger program\n");
    return EXIT_SUCCESS;
}