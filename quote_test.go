package myflyingbox

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testQuoteResp = []byte(`{
      "status":"success",
      "self":"https://test.myflyingbox.com/v2/quotes/ff9abde0-b79a-48b9-a817-cc397d91914e",
      "data":{
        "id":"ff9abde0-b79a-48b9-a817-cc397d91914e",
        "created_at":"2017-03-31T09:47:18+02:00",
        "ordered":true,
        "ordered?":false,
        "shipper":{
          "country":"US",
          "postal_code":"11201",
          "city":"Brooklyn",
          "customer_reference":null
        },
        "recipient":{
          "country":"US",
          "postal_code":"11201",
          "is_a_company":"false",
          "city":"Brooklyn",
          "customer_reference":null
        },
        "parcels":[
          {
            "width":10.0,
            "length":10.0,
            "height":10.0,
            "weight":1.0,
            "value":null,
            "currency":null
          }
        ],
        "offers":[
          {
            "id":"cb2aa075-805a-4d18-8e6a-dfd6862cfa58",
            "quote_id":"ff9abde0-b79a-48b9-a817-cc397d91914e",
            "product_id":"2dd30850-ecb9-4c01-aad1-638b843fe132",
            "orderable?":true,
            "insurable?":false,
            "orderable":true,
            "insurable":false,
            "product":{
              "export_from":[
                "US"
              ],
              "id":"2dd30850-ecb9-4c01-aad1-638b843fe132",
              "logo":"usps",
              "code":"usps_priority_domestic",
              "name":"Priority Domestic",
              "pick_up":false,
              "drop_off":true,
              "preset_delivery_location":false,
              "carrier_code":"usps",
              "delay":"24-72",
              "collection_informations":{
                "en":"Post Office Drop Off",
                "es":"Depósito en oficina de correos USPS",
                "fr":"Dépôt en bureau de poste USPS"
              },
              "delivery_informations":{
                "en":"1 to 3 days",
                "es":"1 a 3 dias",
                "fr":"1 à 3 jours"
              },
              "details":{
                "en":"Delivery at home : YES\r\nGuaranteed delay: NO\r\nInsurance loss MFB : 50€ (invoice requested)\r\nInsurance damage MFB : 0€\r\nInsurance option AD VALOREM : YES \r\nService: from monday to friday",
                "es":"Entrega en oficio: SI\r\nEntrega en casa : SI\r\nPeríodo garantizado : NO\r\nSeguro por perdida MFB : 50€ (con presentacion de factura)\r\nSeguro de daños MFB : 0€\r\nSeguro AD VALOREM : SI \r\nServicio : de lunes a viernes\r\n",
                "fr":"Livraison en bureau : OUI\r\nLivraison à domicile : OUI\r\nDélai Garanti : NON\r\nAssurance Perte MFB : 50€ maximum sur facture\r\nAssurance dommage MFB : 0€\r\nOption Assurance AD VALOREM : OUI \r\nLivraison Relais : NON\r\nService : Lundi au vendredi"
              }
            },
            "price":{
              "formatted":"€7.58",
              "currency":"EUR",
              "amount":"7.58",
              "amount_in_cents":758
            },
            "price_vat":{
              "formatted":"€0.00",
              "currency":"EUR",
              "amount":"0.0",
              "amount_in_cents":0
            },
            "total_price":{
              "formatted":"€7.58",
              "currency":"EUR",
              "amount":"7.58",
              "amount_in_cents":758
            },
            "insurance_price":null,
            "collection_dates":[
              {
                "date":"2017-04-03",
                "cutoff":""
              }
            ]
          },
          {
            "id":"1cb5d716-fd1c-416f-8dfd-dcdba05cc7e2",
            "quote_id":"ff9abde0-b79a-48b9-a817-cc397d91914e",
            "product_id":"631c636f-480a-4398-979b-0e23c5e3a110",
            "orderable?":true,
            "insurable?":false,
            "orderable":true,
            "insurable":false,
            "product":{
              "export_from":[
                "US"
              ],
              "id":"631c636f-480a-4398-979b-0e23c5e3a110",
              "logo":"usps",
              "code":"usps_parcel_select",
              "name":"Parcel Select",
              "pick_up":false,
              "drop_off":true,
              "preset_delivery_location":false,
              "carrier_code":"usps",
              "delay":"48-192",
              "collection_informations":{
                "en":"Post Office Drop Off",
                "es":"Depósito en oficina de correos USPS",
                "fr":"Dépot en bureau de poste USPS"
              },
              "delivery_informations":{
                "en":"2 to 8 days",
                "es":"2 a 8 dias",
                "fr":"2 à 8 jours"
              },
              "details":{
                "en":"Delivery at home : YES\r\nGuaranteed delay: NO\r\nInsurance loss MFB : 50€ (invoice requested)\r\nInsurance damage MFB : 0€\r\nInsurance option AD VALOREM : YES \r\nService: from monday to friday",
                "es":"Entrega en oficio: SI\r\nEntrega en casa : SI\r\nPeríodo garantizado : NO\r\nSeguro por perdida MFB : 50€ (con presentacion de factura)\r\nSeguro de daños MFB : 0€\r\nSeguro AD VALOREM : SI \r\nServicio : de lunes a viernes",
                "fr":"Livraison en bureau : OUI\r\nLivraison à domicile : OUI\r\nDélai Garanti : NON\r\nAssurance Perte MFB : 50€ maximum sur facture\r\nAssurance dommage MFB : 0€\r\nOption Assurance AD VALOREM : OUI \r\nLivraison Relais : NON\r\nService : Lundi au vendredi"
              }
            },
            "price":{
              "formatted":"€7.82",
              "currency":"EUR",
              "amount":"7.82",
              "amount_in_cents":782
            },
            "price_vat":{
              "formatted":"€0.00",
              "currency":"EUR",
              "amount":"0.0",
              "amount_in_cents":0
            },
            "total_price":{
              "formatted":"€7.82",
              "currency":"EUR",
              "amount":"7.82",
              "amount_in_cents":782
            },
            "insurance_price":null,
            "collection_dates":[
              {
                "date":"2017-04-03",
                "cutoff":""
              }
            ]
          },
          {
            "id":"02907cfb-7f58-48a4-8e28-c1a7b4001897",
            "quote_id":"ff9abde0-b79a-48b9-a817-cc397d91914e",
            "product_id":"6922da16-f4e5-42a5-a231-80f9cd1fb1ab",
            "orderable?":true,
            "insurable?":false,
            "orderable":true,
            "insurable":false,
            "product":{
              "export_from":[
                "US",
                "CA"
              ],
              "id":"6922da16-f4e5-42a5-a231-80f9cd1fb1ab",
              "logo":"fedex",
              "code":"fedex_express_saver",
              "name":"Express Saver",
              "pick_up":true,
              "drop_off":false,
              "preset_delivery_location":false,
              "carrier_code":"fedex",
              "delay":"24_72",
              "collection_informations":{
                "en":"Today if you book before 1:00 pm",
                "es":"Recogida hoy por un pedido antes de las 12:00",
                "fr":"Enlèvement le jour même pour toute commande avant 12H (heure locale)"
              },
              "delivery_informations":{
                "en":"1 to 3 days",
                "es":"1 a 3 dias",
                "fr":"1 à 3 jours"
              },
              "details":{
                "en":"Delivery in office : YES\r\nDelivery at home : YES\r\nGuaranteed delay: NO\r\nInsurance loss MFB : 50€ (invoice requested)\r\nInsurance damage MFB : 0€\r\nInsurance option AD VALOREM : YES \r\nService: from monday to friday",
                "es":"Entrega en oficio: SI\r\nEntrega en casa : SI\r\nPeríodo garantizado : NO\r\nSeguro por perdida MFB : 50€ (con presentacion de factura)\r\nSeguro de daños MFB : 0€\r\nSeguro AD VALOREM : SI \r\nServicio : de lunes a viernes",
                "fr":"Livraison à domicile : OUI\r\nDélai Garanti : NON\r\nAssurance Perte MFB : 50€ maximum sur facture\r\nAssurance dommage MFB : 0€\r\nOption Assurance AD VALOREM : OUI \r\nLivraison Relais : NON\r\nService : Lundi au vendredi"
              }
            },
            "price":{
              "formatted":"€11.85",
              "currency":"EUR",
              "amount":"11.85",
              "amount_in_cents":1185
            },
            "price_vat":{
              "formatted":"€0.00",
              "currency":"EUR",
              "amount":"0.0",
              "amount_in_cents":0
            },
            "total_price":{
              "formatted":"€11.85",
              "currency":"EUR",
              "amount":"11.85",
              "amount_in_cents":1185
            },
            "insurance_price":null,
            "collection_dates":[
              {
                "date":"2017-03-31",
                "cutoff":""
              },
              {
                "date":"2017-04-03",
                "cutoff":""
              },
              {
                "date":"2017-04-04",
                "cutoff":""
              }
            ]
          },
          {
            "id":"2a325990-8d23-4aed-a4de-c4140eb22679",
            "quote_id":"ff9abde0-b79a-48b9-a817-cc397d91914e",
            "product_id":"1463b489-adbf-4a0e-ad59-321ff5e2df07",
            "orderable?":true,
            "insurable?":false,
            "orderable":true,
            "insurable":false,
            "product":{
              "export_from":[
                "US",
                "CA"
              ],
              "id":"1463b489-adbf-4a0e-ad59-321ff5e2df07",
              "logo":"fedex",
              "code":"fedex_2_day",
              "name":"2 day",
              "pick_up":true,
              "drop_off":false,
              "preset_delivery_location":false,
              "carrier_code":"fedex",
              "delay":"24 _ 48",
              "collection_informations":{
                "en":"Today if you book before 12:00 am (local time)",
                "es":"Recogida hoy por un pedido antes de las 12:00 (hora local)",
                "fr":"Enlèvement le jour même pour toute commande avant 12H (heure locale)"
              },
              "delivery_informations":{
                "en":"1 - 2 days",
                "es":"1 - 2 dias",
                "fr":"1 - 2 jours"
              },
              "details":{
                "en":"Delivery in office : YES\r\nDelivery at home : YES\r\nGuaranteed delay: NO\r\nInsurance loss MFB : 50€ (invoice requested)\r\nInsurance damage MFB : 0€\r\nInsurance option AD VALOREM : YES \r\nService: from monday to friday",
                "es":"Entrega en oficio: SI\r\nEntrega en casa : SI\r\nPeríodo garantizado : NO\r\nSeguro por perdida MFB : 50€ (con presentacion de factura)\r\nSeguro de daños MFB : 0€\r\nSeguro AD VALOREM : SI \r\nServicio : de lunes a viernes",
                "fr":"Livraison en bureau : OUI\r\nLivraison à domicile : OUI\r\nDélai Garanti : NON\r\nAssurance Perte MFB : 50€ maximum sur facture\r\nAssurance dommage MFB : 0€\r\nOption Assurance AD VALOREM : OUI \r\nLivraison Relais : NON\r\nService : Lundi au vendredi"
              }
            },
            "price":{
              "formatted":"€12.54",
              "currency":"EUR",
              "amount":"12.54",
              "amount_in_cents":1254
            },
            "price_vat":{
              "formatted":"€0.00",
              "currency":"EUR",
              "amount":"0.0",
              "amount_in_cents":0
            },
            "total_price":{
              "formatted":"€12.54",
              "currency":"EUR",
              "amount":"12.54",
              "amount_in_cents":1254
            },
            "insurance_price":null,
            "collection_dates":[
              {
                "date":"2017-03-31",
                "cutoff":""
              },
              {
                "date":"2017-04-03",
                "cutoff":""
              },
              {
                "date":"2017-04-04",
                "cutoff":""
              }
            ]
          },
          {
            "id":"c6e549ff-6671-4d6a-920e-4ef36aa74438",
            "quote_id":"ff9abde0-b79a-48b9-a817-cc397d91914e",
            "product_id":"3c5d6585-c468-4c3d-9867-7b9a69e340c3",
            "orderable?":true,
            "insurable?":false,
            "orderable":true,
            "insurable":false,
            "product":{
              "export_from":[
                "US",
                "CA"
              ],
              "id":"3c5d6585-c468-4c3d-9867-7b9a69e340c3",
              "logo":"fedex",
              "code":"fedex_standard_overnight",
              "name":"Standard Overnight",
              "pick_up":true,
              "drop_off":false,
              "preset_delivery_location":false,
              "carrier_code":"fedex",
              "delay":"24",
              "collection_informations":{
                "en":"Today if you book before 12:00 am",
                "es":"Recogida hoy por un pedido antes de las 12:00",
                "fr":"Enlèvement le jour même pour toute commande avant 12H"
              },
              "delivery_informations":{
                "en":"Next day",
                "es":"Manana",
                "fr":"Le lendemain"
              },
              "details":{
                "en":"Delivery in office : YES\r\nDelivery at home : YES\r\nGuaranteed delay: NO\r\nInsurance loss MFB : 50€ (invoice requested)\r\nInsurance damage MFB : 0€\r\nInsurance option AD VALOREM : YES \r\nService: from monday to friday",
                "es":"Entrega en oficio: SI\r\nEntrega en casa : SI\r\nPeríodo garantizado : NO\r\nSeguro por perdida MFB : 50€ (con presentacion de factura)\r\nSeguro de daños MFB : 0€\r\nSeguro AD VALOREM : SI \r\nServicio : de lunes a viernes",
                "fr":"Livraison en bureau : OUI\r\nLivraison à domicile : OUI\r\nDélai Garanti : NON\r\nAssurance Perte MFB : 50€ maximum sur facture\r\nAssurance dommage MFB : 0€\r\nOption Assurance AD VALOREM : OUI \r\nLivraison Relais : NON\r\nService : Lundi au vendredi"
              }
            },
            "price":{
              "formatted":"€13.65",
              "currency":"EUR",
              "amount":"13.65",
              "amount_in_cents":1365
            },
            "price_vat":{
              "formatted":"€0.00",
              "currency":"EUR",
              "amount":"0.0",
              "amount_in_cents":0
            },
            "total_price":{
              "formatted":"€13.65",
              "currency":"EUR",
              "amount":"13.65",
              "amount_in_cents":1365
            },
            "insurance_price":null,
            "collection_dates":[
              {
                "date":"2017-03-31",
                "cutoff":""
              },
              {
                "date":"2017-04-03",
                "cutoff":""
              },
              {
                "date":"2017-04-04",
                "cutoff":""
              }
            ]
          },
          {
            "id":"7d8b69a1-0715-456f-9904-2c53d3172f71",
            "quote_id":"ff9abde0-b79a-48b9-a817-cc397d91914e",
            "product_id":"5a8b9e21-947e-4215-861d-ba2748c4c822",
            "orderable?":true,
            "insurable?":false,
            "orderable":true,
            "insurable":false,
            "product":{
              "export_from":[
                "US",
                "CA"
              ],
              "id":"5a8b9e21-947e-4215-861d-ba2748c4c822",
              "logo":"fedex",
              "code":"fedex_priority_overnight",
              "name":"Priority Overnight",
              "pick_up":true,
              "drop_off":false,
              "preset_delivery_location":false,
              "carrier_code":"fedex",
              "delay":"24",
              "collection_informations":{
                "en":"Today if you book before 12:00 am (local time)",
                "es":"Recogida hoy por un pedido antes de las 12:00 (hora local)",
                "fr":"Enlèvement le jour même pour toute commande avant 12H (heure local)"
              },
              "delivery_informations":{
                "en":"Next day",
                "es":"Manana",
                "fr":"Le lendemain"
              },
              "details":{
                "en":"Delivery in office : YES\r\nDelivery at home : YES\r\nGuaranteed delay: NO\r\nInsurance loss MFB : 50€ (invoice requested)\r\nInsurance damage MFB : 0€\r\nInsurance option AD VALOREM : YES \r\nService: from monday to saturday",
                "es":"Entrega en oficio: SI\r\nEntrega en casa : SI\r\nPeríodo garantizado : NO\r\nSeguro por perdida MFB : 50€ (con presentacion de factura)\r\nSeguro de daños MFB : 0€\r\nSeguro AD VALOREM : SI \r\nServicio : de lunes a sábado",
                "fr":"Livraison en bureau : OUI\r\nLivraison à domicile : OUI\r\nDélai Garanti : NON\r\nAssurance Perte MFB : 50€ maximum sur facture\r\nAssurance dommage MFB : 0€\r\nOption Assurance AD VALOREM : OUI \r\nLivraison Relais : NON\r\nService : Lundi au samedi"
              }
            },
            "price":{
              "formatted":"€14.28",
              "currency":"EUR",
              "amount":"14.28",
              "amount_in_cents":1428
            },
            "price_vat":{
              "formatted":"€0.00",
              "currency":"EUR",
              "amount":"0.0",
              "amount_in_cents":0
            },
            "total_price":{
              "formatted":"€14.28",
              "currency":"EUR",
              "amount":"14.28",
              "amount_in_cents":1428
            },
            "insurance_price":null,
            "collection_dates":[
              {
                "date":"2017-03-31",
                "cutoff":""
              },
              {
                "date":"2017-04-03",
                "cutoff":""
              },
              {
                "date":"2017-04-04",
                "cutoff":""
              }
            ]
          },
          {
            "id":"17b8761a-a22e-4a93-b7e2-113772279239",
            "quote_id":"ff9abde0-b79a-48b9-a817-cc397d91914e",
            "product_id":"2fe5469d-8849-4814-8484-91bff6983c37",
            "orderable?":true,
            "insurable?":false,
            "orderable":true,
            "insurable":false,
            "product":{
              "export_from":[
                "US"
              ],
              "id":"2fe5469d-8849-4814-8484-91bff6983c37",
              "logo":"usps",
              "code":"usps_priority_express",
              "name":"Priority Express",
              "pick_up":false,
              "drop_off":true,
              "preset_delivery_location":false,
              "carrier_code":"usps",
              "delay":"24-48",
              "collection_informations":{
                "en":"Post Office Drop Off",
                "es":"Depósito en oficina de correos USPS",
                "fr":"Dépôt en bureau de poste USPS"
              },
              "delivery_informations":{
                "en":"1 to 2 days",
                "es":"1 a 2 dias",
                "fr":"1 à 2 jours"
              },
              "details":{
                "en":"Delivery at home : YES\r\nGuaranteed delay: NO\r\nInsurance loss MFB : 50€ (invoice requested)\r\nInsurance damage MFB : 0€\r\nInsurance option AD VALOREM : YES \r\nService: from monday to friday",
                "es":"Entrega en oficio: SI\r\nEntrega en casa : SI\r\nPeríodo garantizado : NO\r\nSeguro por perdida MFB : 50€ (con presentacion de factura)\r\nSeguro de daños MFB : 0€\r\nSeguro AD VALOREM : SI \r\nServicio : de lunes a viernes",
                "fr":"Livraison en bureau : OUI\r\nLivraison à domicile : OUI\r\nDélai Garanti : NON\r\nAssurance Perte MFB : 50€ maximum sur facture\r\nAssurance dommage MFB : 0€\r\nOption Assurance AD VALOREM : OUI \r\nLivraison Relais : NON\r\nService : Lundi au vendredi"
              }
            },
            "price":{
              "formatted":"€27.31",
              "currency":"EUR",
              "amount":"27.31",
              "amount_in_cents":2731
            },
            "price_vat":{
              "formatted":"€0.00",
              "currency":"EUR",
              "amount":"0.0",
              "amount_in_cents":0
            },
            "total_price":{
              "formatted":"€27.31",
              "currency":"EUR",
              "amount":"27.31",
              "amount_in_cents":2731
            },
            "insurance_price":null,
            "collection_dates":[
              {
                "date":"2017-04-03",
                "cutoff":""
              }
            ]
          },
          {
            "id":"4acc77e0-2ac9-4bcd-9d03-6960c7b3c567",
            "quote_id":"ff9abde0-b79a-48b9-a817-cc397d91914e",
            "product_id":"f989f545-cd0e-4214-b225-322aa157b372",
            "orderable?":true,
            "insurable?":false,
            "orderable":true,
            "insurable":false,
            "product":{
              "export_from":[
                "US",
                "CA"
              ],
              "id":"f989f545-cd0e-4214-b225-322aa157b372",
              "logo":"dhl",
              "code":"dhl_domestic_express_18",
              "name":"Express 18",
              "pick_up":true,
              "drop_off":false,
              "preset_delivery_location":false,
              "carrier_code":"dhl",
              "delay":"24",
              "collection_informations":{
                "en":"Today if you book before 12:00 am",
                "es":"Recogida hoy por un pedido antes de las 12:00",
                "fr":"Enlèvement le jour même pour toute commande avant 12H"
              },
              "delivery_informations":{
                "en":"Next day*",
                "es":"Manana*",
                "fr":"Livraison le lendemain*"
              },
              "details":{
                "en":"Delivery in office : YES\r\nDelivery at home : YES\r\nGuaranteed delay: NO\r\nInsurance loss MFB : 50€ (invoice requested)\r\nInsurance damage MFB : 0€\r\nInsurance option AD VALOREM : YES \r\nService: from monday to friday",
                "es":"Entrega en oficio: SI\r\nEntrega en casa : SI\r\nPeríodo garantizado : NO\r\nSeguro por perdida MFB : 50€ (con presentacion de factura)\r\nSeguro de daños MFB : 0€\r\nSeguro AD VALOREM : SI \r\nServicio : de lunes a viernes",
                "fr":"Livraison en bureau : OUI\r\nLivraison à domicile : OUI\r\nDélai Garanti : NON\r\nAssurance Perte MFB : 50€ maximum sur facture\r\nAssurance dommage MFB : 0€\r\nOption Assurance AD VALOREM : OUI \r\nLivraison Relais : NON\r\nService : Lundi au vendredi"
              }
            },
            "price":{
              "formatted":"€35.35",
              "currency":"EUR",
              "amount":"35.35",
              "amount_in_cents":3535
            },
            "price_vat":{
              "formatted":"€0.00",
              "currency":"EUR",
              "amount":"0.0",
              "amount_in_cents":0
            },
            "total_price":{
              "formatted":"€35.35",
              "currency":"EUR",
              "amount":"35.35",
              "amount_in_cents":3535
            },
            "insurance_price":null,
            "collection_dates":[
              {
                "date":"2017-03-31",
                "cutoff":""
              },
              {
                "date":"2017-04-03",
                "cutoff":""
              },
              {
                "date":"2017-04-04",
                "cutoff":""
              }
            ]
          },
          {
            "id":"8fa2b874-4bfe-4ab1-bf0a-881ee90ffd78",
            "quote_id":"ff9abde0-b79a-48b9-a817-cc397d91914e",
            "product_id":"4d665a9c-9612-467d-be66-9c1a98d04fed",
            "orderable?":true,
            "insurable?":false,
            "orderable":true,
            "insurable":false,
            "product":{
              "export_from":[
                "US",
                "CA"
              ],
              "id":"4d665a9c-9612-467d-be66-9c1a98d04fed",
              "logo":"fedex",
              "code":"fedex_first_overnight",
              "name":"First Overnight",
              "pick_up":true,
              "drop_off":false,
              "preset_delivery_location":false,
              "carrier_code":"fedex",
              "delay":"24",
              "collection_informations":{
                "en":"Today if you book before 12:00 am",
                "es":"Recogida hoy por un pedido antes de las 12:00",
                "fr":"Enlèvement le jour même pour toute commande avant 12H"
              },
              "delivery_informations":{
                "en":"Next day before 10:00",
                "es":"Manana antes de las 10:00",
                "fr":"Le lendemain avant 10:00\r\n"
              },
              "details":{
                "en":"Delivery in office : YES\r\nDelivery at home : YES\r\nGuaranteed delay: NO\r\nInsurance loss MFB : 50€ (invoice requested)\r\nInsurance damage MFB : 0€\r\nInsurance option AD VALOREM : YES \r\nService: from monday to friday",
                "es":"Entrega en oficio: SI\r\nEntrega en casa : SI\r\nPeríodo garantizado : NO\r\nSeguro por perdida MFB : 50€ (con presentacion de factura)\r\nSeguro de daños MFB : 0€\r\nSeguro AD VALOREM : SI \r\nServicio : de lunes a viernes",
                "fr":"Livraison en bureau : OUI\r\nLivraison à domicile : OUI\r\nDélai Garanti : NON\r\nAssurance Perte MFB : 50€ maximum sur facture\r\nAssurance dommage MFB : 0€\r\nOption Assurance AD VALOREM : OUI \r\nLivraison Relais : NON\r\nService : Lundi au vendredi"
              }
            },
            "price":{
              "formatted":"€78.52",
              "currency":"EUR",
              "amount":"78.52",
              "amount_in_cents":7852
            },
            "price_vat":{
              "formatted":"€0.00",
              "currency":"EUR",
              "amount":"0.0",
              "amount_in_cents":0
            },
            "total_price":{
              "formatted":"€78.52",
              "currency":"EUR",
              "amount":"78.52",
              "amount_in_cents":7852
            },
            "insurance_price":null,
            "collection_dates":[
              {
                "date":"2017-03-31",
                "cutoff":""
              },
              {
                "date":"2017-04-03",
                "cutoff":""
              },
              {
                "date":"2017-04-04",
                "cutoff":""
              }
            ]
          }
        ],
        "origin":null
      }
    }`)
)

func TestParseQuote(t *testing.T) {

	var apiResp Response
	var q Quote
	assert.NoError(t, json.Unmarshal(testQuoteResp, &apiResp))
	assert.NoError(t, apiResp.Error())

	data, err := json.Marshal(apiResp.Data)
	assert.NoError(t, err)
	assert.NoError(t, json.Unmarshal(data, &q))
	assert.True(t, q.Ordered)
	assert.Equal(t, "US", q.Shipper.Country)
	assert.Equal(t, "Brooklyn", q.Recipient.City)
	assert.Equal(t, 1, len(q.Parcels))
	assert.Equal(t, float64(10), q.Parcels[0].Width)
	if !assert.True(t, len(q.Offers) > 1) {
		return
	}
	assert.Len(t, q.Offers, 9)
	assert.Equal(t, "2dd30850-ecb9-4c01-aad1-638b843fe132", q.Offers[0].Product.ID)
	assert.Equal(t, "Post Office Drop Off", q.Offers[0].Product.CollectionInformations["en"])
	assert.Equal(t, "1 à 3 jours", q.Offers[0].Product.DeliveryInformations["fr"])
	assert.Equal(t, 7.58, q.Offers[0].Price.Amount)
	assert.Equal(t, "2017-04-03", q.Offers[0].CollectionDates[0].Date)
	assert.Nil(t, q.Offers[0].InsurancePrice)
}

func TestGetQuoteID(t *testing.T) {
	quoteID, err := getQuoteID("foobar")
	assert.NoError(t, err)
	assert.Equal(t, "foobar", quoteID)

	quoteID, err = getQuoteID(&quoteID)
	assert.NoError(t, err)
	assert.Equal(t, "foobar", quoteID)

	quoteID, err = getQuoteID(Quote{ID: "foobar"})
	assert.NoError(t, err)
	assert.Equal(t, "foobar", quoteID)

	quoteID, err = getQuoteID(&Quote{ID: "foobar"})
	assert.NoError(t, err)
	assert.Equal(t, "foobar", quoteID)

	_, err = getQuoteID(&Order{})
	assert.Equal(t, ErrInvalidArgumentType, err)
}

func TestPlaceQuote(t *testing.T) {
	api := getAPI()
	q := Quote{
		Shipper:   Shipper{Country: "US", PostalCode: "11201", City: "Brooklyn"},
		Recipient: Recipient{IsACompany: false, Country: "US", PostalCode: "11201", City: "Brooklyn"},
		Parcels: []Parcel{
			Parcel{Weight: 1.0, Length: 10, Width: 10, Height: 10},
		},
	}
	res, err := api.RequestQuote(ctx, &q)
	assert.NoError(t, err)
	if !assert.NotNil(t, res) {
		return
	}
	// ID's are uuid, so expect a length of 36
	if !assert.Len(t, res.ID, 36) {
		return
	}
	res2, err := api.RetrieveQuote(ctx, res)
	assert.NoError(t, err)
	assert.Equal(t, res.ID, res2.ID)
	if !assert.True(t, len(res2.Offers) > 0) {
		t.Logf("Retrieved quote: %#v", res2)
		return
	}
}
