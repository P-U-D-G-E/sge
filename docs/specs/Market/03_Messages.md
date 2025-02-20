# **Messages**

The Market module exposes the following services:

- AddMarket
- ResolveMarket
- UpdateMarket

```proto
// Msg defines the Msg service.
service Msg {
    rpc Add(MsgAdd) returns (MarketResponse);
    rpc Resolve(MsgResolve) returns (MarketResponse);
    rpc Update(MsgUpdate) returns (MarketResponse);
}
```

---

## **MsgAdd**

This message is used to add new market to the chain

```proto
message MsgAdd {
  string creator = 1;
  string ticket = 2;
}
```

### Add Market Ticked Payload

The ticket data for market addition is as follows:

```proto
// MarketAddTicketPayload indicates data of add market ticket
message MarketAddTicketPayload {
  // uid is the universal unique identifier of the market.
  string uid = 1 [
    (gogoproto.customname) = "UID",
    (gogoproto.jsontag) = "uid",
    json_name = "uid"
  ];
  // start_ts is the start timestamp of the market.
  uint64 start_ts = 2 [
    (gogoproto.customname) = "StartTS",
    (gogoproto.jsontag) = "start_ts",
    json_name = "start_ts"
  ];
  // end_ts is the end timestamp of the market.
  uint64 end_ts = 3 [
    (gogoproto.customname) = "EndTS",
    (gogoproto.jsontag) = "end_ts",
    json_name = "end_ts"
  ];
  // odds is the list of odds of the market.
  repeated Odds odds = 4;

  // status is the current status of the market.
  MarketStatus status = 5;

  // creator is the address of the creator of the market.
  string creator = 6;

  // meta contains human-readable metadata of the market.
  string meta = 7;
}
```

#### **Sample addition ticket**

```json
{
    "uid": "5531c60f-2025-48ce-ae79-1dc110f16000",
    "start_ts": 1668480139,
    "end_ts": 1883781609,
    "odds": [
        {
            "uid": "9991c60f-2025-48ce-ae79-1dc110f16990",
            "meta": "x is winner"
        },
        {
            "uid": "9991c60f-2025-48ce-ae79-1dc110f16991",
            "meta": "y is winner"
        },
        {
            "uid": "9991c60f-2025-48ce-ae79-1dc110f16992",
            "meta": "draw"
        }
    ],
    "status": 1,
    "meta": "Soccer: England vs USA",
    "iat": 1665140310,
    "exp": 1757788212
}
```

---

## **MsgUpdate**

This message is used to update already existent markets on the chain

```proto
// MsgUpdate is the message type for updating market data.
// in the state
message MsgUpdate {
  // creator is the address of the creator account of the market.
  string creator = 1;
  // ticket is the jwt ticket data.
  string ticket = 2;
}
```

### Update Market Ticked Payload

```proto
// MarketUpdateTicketPayload indicates data of the market update ticket
message MarketUpdateTicketPayload {
  // uid is the uuid of the market
  string uid = 1 [
    (gogoproto.customname) = "UID",
    (gogoproto.jsontag) = "uid",
    json_name = "uid"
  ];
  // start_ts is the start timestamp of the market
  uint64 start_ts = 2 [
    (gogoproto.customname) = "StartTS",
    (gogoproto.jsontag) = "start_ts",
    json_name = "start_ts"
  ];
  // end_ts is the end timestamp of the market
  uint64 end_ts = 3 [
    (gogoproto.customname) = "EndTS",
    (gogoproto.jsontag) = "end_ts",
    json_name = "end_ts"
  ];
  // status is the status of the resolution.
  MarketStatus status = 4;
}
```

#### **Sample update ticket**

```json
{
    "uid": "5531c60f-2025-48ce-ae79-1dc110f16000",
    "start_ts": 1668480139,
    "end_ts": 1883781609,
    "status": 1,
    "iat": 1665140310,
    "exp": 1757788212
}
```

---

## **MarketResponse**

This is the common response to all the messages

```proto
// MsgAddResponse response for adding market.
message MsgAddResponse {
  // error contains an error if adding a market faces any issues.
  string error = 1 [ (gogoproto.nullable) = true ];
  // data is the data of market.
  Market data = 2 [ (gogoproto.nullable) = true ];
}
```

---

## **MsgResolve**

This message is used to resolve already existent markets on the chain

```proto
// MsgResolve is the message type for resolving a market.
message MsgResolve {
  // creator is the address of the creator account of the market.
  string creator = 1;
  // ticket is the jwt ticket data.
  string ticket = 2;
}
```

### Resolve Market Ticked Payload

```proto
// MarketResolutionTicketPayload indicates data of the
// resolution of the market ticket.
message MarketResolutionTicketPayload {
  // uid is the universal unique identifier of the market.
  string uid = 1 [
    (gogoproto.customname) = "UID",
    (gogoproto.jsontag) = "uid",
    json_name = "uid"
  ];

  // resolution_ts is the resolution timestamp of the market.
  uint64 resolution_ts = 2 [
    (gogoproto.customname) = "ResolutionTS",
    (gogoproto.jsontag) = "resolution_ts",
    json_name = "resolution_ts"
  ];

  // winner_odds_uids is the universal unique identifier list of the winner
  // odds.
  repeated string winner_odds_uids = 3 [
    (gogoproto.customname) = "WinnerOddsUIDs",
    (gogoproto.jsontag) = "winner_odds_uids",
    json_name = "winner_odds_uids"
  ];

  // status is the status of the resolution.
  MarketStatus status = 4;
}
```

#### **Sample resolve ticket**

```json
{
    "uid": "5531c60f-2025-48ce-ae79-1dc110f16000",
    "resolution_ts": 1668480139,
    "winner_odds_uids": [
      "9991c60f-2025-48ce-ae79-1dc110f16990"
    ],
    "status": 1,
    "iat": 1665140310,
    "exp": 1757788212
}
```
