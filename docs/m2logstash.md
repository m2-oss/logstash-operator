# M2Logstash
CRD определяет параметры запуска инстансов M2Logstash. Инстансы запускаются, как StatefulSet в namespace, где установлена CRD

## M2LogstashSpec
`image (string, optional)` Default: 'logstash'\
`tag (string, optional)` Default: '7.16.2'\
`name (string, optional)` Default: 'logstash'. Название инстанса\
`java_opts (string, optional)` Default: '-Xmx1g -Xms1g'. Параметры запуска jvm\
`resources (M2LogstashSpecResources, optional)`

## M2LogstashSpecResources
`limits (M2LogstashSpecLimit, optional)`
`requests (M2LogstashSpecRequests, optional)`

## M2LogstashSpecRequests
`cpu (string, optional)` Default: '100m'
`memory (string, optional)` Default '1536Mi'

## M2LogstashSpecLimit
`cpu (string, optional)` Default: '2000m'
`memory (string, optional)` Default '1536Mi'