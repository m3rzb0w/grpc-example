import grpc
import products_pb2
import products_pb2_grpc

def check_product_availability(product_id, quantity):
    channel = grpc.insecure_channel('localhost:50051')
    stub = products_pb2_grpc.ProductsServiceStub(channel)

    request = products_pb2.ProductRequest(
        product_id=product_id,
        quantity=quantity
    )
    
    try:
        response = stub.CheckProductsAvailability(request)
        
        return {
            "available": response.available,
            "ean": response.ean,
            "sku": response.sku,
            "brut_price": response.brut_price,
            "net_price": response.net_price,
            "additional_info": response.additional_info
        }
    except grpc.RpcError as e:
        return {"error": str(e)}

if __name__ == '__main__':
    product_id = "PROD001"
    quantity = 1
    result = check_product_availability(product_id, quantity)
    print(result)

    product_id = "PROD002"
    quantity = 3
    result = check_product_availability(product_id, quantity)
    print(result)

    product_id = "PROD003"
    quantity = 20
    result = check_product_availability(product_id, quantity)
    print(result)

    product_id = "PROD004"
    quantity = 2
    result = check_product_availability(product_id, quantity)
    print(result)

    product_id = "PROD005"
    quantity = 200
    result = check_product_availability(product_id, quantity)
    print(result)

    product_id = "PROD006"
    quantity = 5
    result = check_product_availability(product_id, quantity)
    print(result)