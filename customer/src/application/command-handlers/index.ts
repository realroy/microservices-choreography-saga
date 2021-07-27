import { CustomerRepository } from ".."
import { CreditLimitExceeded, CreditReserved, Customer } from "../../core"

export const reserveCredit =  async (amount: number, customerId: string, customerRepository: CustomerRepository) => {
  let customer: Customer
  
  try {
    customer = await customerRepository.findOne({ id: customerId })
  } catch (error) {
    throw error
  }

  if (customer.credit < amount) {
    return new CreditLimitExceeded({ customer })
  }

  try {
    customer = await customerRepository.updateOne({
      where: { id: customer.id },
      with: { credit: customer.credit - amount }
    })
  } catch (error) {
    throw error
  }

  return new CreditReserved({ customer })
}
