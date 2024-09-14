Buy Now Pay Later (BNPL) Feature
Overview
This project implements a simple Buy Now Pay Later (BNPL) system for an e-commerce platform in Go. The system includes key functionalities such as inventory management, order placement with payment options, user credit tracking, and user blacklisting for defaulted payments.

Features
Inventory Management:
- Add and update products in the inventory.
- View the current inventory status.

Order Placement:
- Users can place orders for available products using two payment modes:
- Prepaid: Immediate payment and order completion.
- BNPL: Buy now, pay within 30 days (subject to credit and blacklist status).

Payment Management:
- Clear pending dues either partially or completely.
- User Blacklisting:
- Users are blacklisted if they default on 3 orders. Blacklisted users cannot place further orders using BNPL.


git clone https://github.com/GHATAK123/bnpl.git
cd bnpl

To Run : go run main.go
