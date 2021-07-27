import { Customer } from "../../core";

export interface CustomerRepository {
  findMany(conditions?: Record<string, any>): Promise<Customer[]>
  findOne(conditions?: Record<string, any>): Promise<Customer>
  create(args: Partial<Customer>): Promise<Customer>
  updateOne(args: { where: Record<string, any>, with: Partial<Customer> }): Promise<Customer>

}