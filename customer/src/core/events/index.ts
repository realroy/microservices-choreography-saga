import { Customer } from '../models'

class Event<T> {
  constructor(args: Partial<T>) {
    Object.assign(this, args)
  }
}

export class CreditReserved extends Event<CreditReserved> {
  readonly customer!: Customer
  readonly orderID!: string
}
export class CreditLimitExceeded extends Event<CreditLimitExceeded> {
  readonly customer!: Customer
  readonly orderId!: string
} 