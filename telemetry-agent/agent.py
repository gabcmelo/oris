import hashlib
import json
import os
import random
import time
import urllib.request

OTLP_ENDPOINT = os.getenv("OTLP_ENDPOINT", "http://otel-collector:4318/v1/metrics")
ENABLED = os.getenv("TELEMETRY_ENABLED", "false").lower() == "true"
SEED = os.getenv("INSTANCE_ID_SEED", "local")

instance_id = hashlib.sha256((SEED + str(int(time.time() / 86400))).encode()).hexdigest()[:16]


while True:
    if ENABLED:
        payload = {
            "resourceMetrics": [
                {
                    "resource": {
                        "attributes": [
                            {"key": "service.name", "value": {"stringValue": "safeguild-telemetry-agent"}},
                            {"key": "instance.id", "value": {"stringValue": instance_id}},
                        ]
                    },
                    "scopeMetrics": [
                        {
                            "metrics": [
                                {
                                    "name": "safeguild.cpu.utilization",
                                    "gauge": {"dataPoints": [{"asDouble": random.random()}]},
                                },
                                {
                                    "name": "safeguild.api.latency.p95",
                                    "gauge": {"dataPoints": [{"asDouble": random.uniform(10, 120)}]},
                                },
                            ]
                        }
                    ],
                }
            ]
        }

        req = urllib.request.Request(
            OTLP_ENDPOINT,
            data=json.dumps(payload).encode("utf-8"),
            headers={"Content-Type": "application/json"},
            method="POST",
        )
        try:
            urllib.request.urlopen(req, timeout=3)
        except Exception:
            pass
    time.sleep(15)
