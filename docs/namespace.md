# Namespace labels and annotations
На конфигурацию инстансов logstash влияют также ряд лейблов и аннотаций, установленных на namespace

Labels:
- `m2/logger (bool, optional)` - определяет нужно ли создавать пулл Logstash для данного namespace
- `m2/logger-ilm-policy (string, optional)` - название ILM Policy в случае использования в качестве output Elasticsearch
- `m2/logger-type` (string, optional)` - префикт для названия индексов в случае использования в качестве output Elasticsearch

Annotations:
- `logger.m2.ru/logstash-replicas (int, optional)` число реплик пуле Logstash, минимальное значение 1optional
___
# Namespace labels and annotations

