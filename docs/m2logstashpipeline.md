# M2LogstashPipeline
CRD определяет конфигурации цепочек input, filter и output.

## M2LogstashPipelineSpec
`input (M2LogstashInputSpec, required)`\
`output (M2LogstashOutputSpec, required)`\
`filter (string, optional)`\

## M2LogstashInputSpec
`kafka (M2LogstashInputSpecKafka, optional)`\
`beats (M2LogstashInputSpecBeats, optional)`\
`http (M2LogstashInputSpecHTTP, optional)`\
`s3 (M2LogstashInputSpecS3, optional)`\
`udp (M2LogstashInputSpecUDP, optional)`\
`tcp (M2LogstashInputSpecTCP, optional)`\
`dlq (M2LogstashInputSpecDLQ, optional)`\

## M2LogstashOutputSpec
`elasticsearch (M2LogstashOutputSpecElasticsearch, optional)`\
`graphite (M2LogstashOutputGraphite, optional)`\
`s3 (M2LogstashOutputSpecS3, optional)`\
`udp (M2LogstashOutputSpecUDP, optional)`\
`tcp (M2LogstashOutputSpecTCP, optional)`\

## M2LogstashInputSpecKafka
`id (string, optional)` Default: 'kafka'\
`hosts ([]string, required)` Default: -\
`decorate_events (string, optional)` Default: 'none'\
`auto_offset_reset (string, optional)` Default: 'earlest'\
`consumer_threads (int, optional)` Default: 1\
`security_protocol (string, optional)` Default: 'PLAINTEXT'\
`sasl_mechanism (string, optional)` Default: 'GSSAPI'\
`secret (string, optional)` Default: - Названия ресурса Secret, где хранятся KAFKA_USERNAME, KAFKA_PASSWORD
`topic` (string, optional) Если не задан, название будет формироваться по названию namespace, которому привязн Logstash

## M2LogstashInputSpecBeats
`id (string, optional)` Default: 'beats'\
`port (int, optional)` Default: -\

## M2LogstashInputSpecHTTP
`id (string, optional)` Default: 'http'\
`port (int, optional)` Default: -\
`secret (string, optional)` Названия ресурса Secret, где хранятся HTTP_USERNAME, HTTP_PASSWORD

## M2LogstashInputSpecTCP
`id (string, optional)` Default: 'tcp'\
`port (int, optional)` Default: -\

## M2LogstashInputSpecUDP
`id (string, optional)` Default: 'udp'\
`port (int, optional)` Default: -\
`queue_size (int, optional)` Default: -\

##M2LogstashInputSpecS3
`id (string, optional)` Default: 's3'\
`bucket (string, optional)` В случае пустого значения, название будет формировать из [имени CRD]-[namespace]\
`endpoint (string, optional)` Default: -\
`region (string, optional)` Default: 'us-east-1'\
`gzip_pattern (string, optional)` Default: '\\.gz(ip)?$'\
`delete (bool, optional)` Default: -\
`exclude_pattern (string, optional)` Default: -\
`secret (string, required)` Названия ресурса Secret, где хранятся ACCESS_KEY_IDб SECRET_ACCESS_KEY\
`prefix (string, optional)` Default: -\
``

## M2LogstashInputSpecDLQ
`id (string, optional)` Default: 'dlq'\
`path (string, required)` Default: -\
`commit_offsets (string, optional)` Default: -\
`pipeline_id (string, optional)` Default: 'main'\

## M2LogstashOutputSpecElasticsearch
`hosts ([]string, required)` Default: -\
`ssl (bool, optional)` Default: -\
`ssl_certificate_verification (bool, optional)` Default: -\
`cacert (string, optional)` Путь внутри контейнера, куда будет смонтирован CA сертификат из Secret logstash-elasticsearch-ca-[CRD name]\
`secret (string, optional)` Названия ресурса Secret, где хранятся ELASTIC_USERNAME, ELASTIC_PASSWORD\
`ilm (bool, optional)` Default: -\
`index (string, optional)` Если индекс не задан, название будет формироваться по названию namespace, которому привязн Logstash

## M2LogstashOutputGraphite
`host (string, required)` Default: -\
`port (int, optional)` Default: -\
`reconnect_interval (int, optional)` Default: 2\
`resend_on_failure (bool, optional)` Default: -\
`timestamp_field (string, optional)` Default: '@timestamp'\

## M2LogstashOutputSpecS3
`bucket (string, optional)` В случае пустого значения, название будет формировать из [имени CRD]-[namespace]\
`canned_acl (string, optional)` Default: 'private'\
`encoding (string, optional)` Default: 'none'\
`endpoint` (string, optional) Default: -\
`region (string, optional)` Default: 'us-east-1'\
`rotation_strategy (string, optional)` Default: 'size_and_time'\
`size_file (int, optional)` Default: 5242880\
`time_file (int, optional)` Default: 15\
`upload_worker_count, (int, optional)` Default: 4\
`secret (string, required)` Названия ресурса Secret, где хранятся ACCESS_KEY_IDб SECRET_ACCESS_KEY\

## M2LogstashOutputSpecUDP
`host (string, required)` Default: -\
`port (int, required)` Default: -\

## M2LogstashOutputSpecUDP
`host (string, required)` Default: -\
`port (int, required)` Default: -\
