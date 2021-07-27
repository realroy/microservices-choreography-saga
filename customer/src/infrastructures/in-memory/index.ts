import { CustomerRepository } from "../../application"
import { Customer } from "../../core"

const customers: Record<string, Customer> = {
  "1": {
    id: "1",
    credit: 100,
  }
}

export class CustomerInMemoryRepository implements CustomerRepository {
  async findMany(conditions?: Record<string, any>): Promise<Customer[]> {
    return Object.values(customers)
  }

  async create(args: Partial<Customer>): Promise<Customer> {
    const customer = new Customer()
    customer.credit = args.credit || 0

    customer.id = new Array(16).fill(0).map(() => Math.round(Math.random() * 9)).join()
  
    return customer
  }

  async findOne(conditions?: Record<string, any>): Promise<Customer> {
    console.log({ conditions })
    const id = conditions?.id
    
    return customers[id]
  }

  async updateOne(args: { where: Record<string, any>; with: Partial<Customer> }): Promise<Customer> {
    const id = args.where?.id
    console.log({ id })

    const customer = customers[id]

    if (!customer) {
      throw new Error('resource not found')
    }

    customers[id] = { ...customer, ...args.with }

    return customers[id]
  }

} 