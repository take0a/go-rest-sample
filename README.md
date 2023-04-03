# Golang による RDBMS をデータソースとする REST API のサンプル

## 外部仕様
### データベース

- 以下の４テーブルを対象とする

```mermaid
erDiagram
    CUSTOMER ||..o{ ORDER : "注文する"
    CUSTOMER {
        int customerId PK "顧客ID"
        string name "氏名"
        string address "住所"
    }
    ORDER ||--|{ ORDER_DETAIL : "構成する"
    ORDER {
        int orderId PK "受注ID"
        int customerId FK "顧客ID"
        date orderDate "受注日"
    }
    ORDER_DETAIL {
        int orderId PK "受注ID"
        int rowNumber PK "行番号"
        int productID FK "製品ID"
        int quantity "数量"
        int pricePerUnit "販売単価"
    }
    ORDER_DETAIL }o..|| PRODUCT : "販売する"
    PRODUCT {
        int productID PK "製品ID"
        string name "製品名"
        int pricePerUnit "標準単価"
    }
```

### REST リソース

- 従属エンティティである ORDER_DETAIL は 独立エンティティである ORDER に吸収させて３リソースにする

```mermaid
erDiagram
    CUSTOMER {
        int customerId PK "顧客ID"
        string name "氏名"
        string address "住所"
    }
    ORDER ||--|{ ORDER_DETAIL : "構成する"
    ORDER {
        int orderId PK "受注ID"
        int customerId FK "顧客ID"
        date orderDate "受注日"
    }
    ORDER_DETAIL {
        int orderId PK "受注ID"
        int rowNumber PK "行番号"
        int productID FK "製品ID"
        int quantity "数量"
        int pricePerUnit "販売単価"
    }
    PRODUCT {
        int productID PK "製品ID"
        string name "製品名"
        int pricePerUnit "標準単価"
    }
```
