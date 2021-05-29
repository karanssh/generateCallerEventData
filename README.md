# generateCallerEventData

generates random data as per following spec:

```json
{
    "event_source": "text",
    "event_ref" : "uuid",
    "event_type": 1,
    "event_date": "time.Time",
    "calling_number": 123456789,
    "called_number": 123456789,
    "location": "kolkata",
    "duration_seconds": 20,
    "attr_1": "text",
    "attr_2": "text",
    "attr_3": "text",
    "attr_4": "text",
    "attr_5": "text",
    "attr_6": "text",
    "attr_7": "text",
    "attr_8": "text"
}
```

## How to use the data loader

All the binary settings can be controlled using config json file located in config.json file. Following parameters can be changed: 

```json
{
    "OutputFormat": "csv",
    "DataSetSize": 100000,
    "DataSetBreaker": 10000,
    "Type1": 15,
    "Type2": 20,
    "Type3": 20,
    "Type4": 45,
    "RandomTextSize": 7,
    "PhoneNoLow": 923152265,
    "PhoneNoHigh": 980809969,
    "DurationLow": 5,
    "DurationHigh": 900
}

```
