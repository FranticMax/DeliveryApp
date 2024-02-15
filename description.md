Курьерская служба

**OrdersService - создание заказов на доставку, оплата заказов**
    Контракты:
    - [POST] /order {"recipient", "shipment"}
    - [POST] /order/pay - редирект на сервис оплаты
    - [DELETE] /order/{id}
    - [GET] /order/{id}

    События:
    - Порождает ORDER_CREATED
    - Порождает ORDER_CANCELED

**DeliveryService - управление движением заказов**
    Контракты:
    - [GET] /delivery/{id} - получить информацию о доставке
    - [GET] /delivery/ - получить информацию о всех доставках
    События:
    - Слушает ORDER_CREATED
    - Слушает ORDER_CANCELED
    - Публикает ORDER_DELIVERED (вручен получателю)
    - Публикает ORDER_TRANSPORTED (доставлен в сортировочный центр)
    - Слушает OREDER_PREPARED_TO_DELIVERY
    - Слушает OREDER_PREPARED_TO_TRANSPORTING

**SortingService - сортировка/подготовка заказов к дальнейшем транспортировке**
    Контракты:
    - [GET] /sorting/transporting/{id} - получить информацию о транспортировке

    События:
    - Слушает ORDER_TRANSPORTED
    - Публикает OREDER_PREPARED_TO_DELIVERY
    - Публикает OREDER_PREPARED_TO_TRANSPORTING

**MonitoringService - мониторинг статусов заказа**
    Контракты:
    - [GET] /monitoring/order/{id} - получить статус по заказу
    - [GET] /monitoring/client/{id} - получить стату по всем заказам клиента

    События:
    - Слушает все события