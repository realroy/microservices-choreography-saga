import express from 'express'
import redis from 'redis'
import { reserveCredit } from '../../application'
import { CreditLimitExceeded } from '../../core'
import { CustomerInMemoryRepository } from '../../infrastructures'


const PORT = process.env.POST || 3001

function start() {
  const app = express()
  const customerRepository = new CustomerInMemoryRepository()

  const subscriber = redis.createClient()
  const publisher = redis.createClient()

  subscriber.subscribe("order:created")

  subscriber.on("message", async (channel, message) => {
    console.log("Received data from channel: %s\n%s", channel, message)
    const parsedMessage = JSON.parse(message)
    switch (channel) {
      case 'order:created':
        const result = await reserveCredit(parsedMessage["amount"], parsedMessage["customer_id"], customerRepository)
        const event = JSON.stringify({ customer_id: result.customer.id, order_id: parsedMessage["id"] })
        
        if (result instanceof CreditLimitExceeded) {
          console.log('credit:limit_exceeded', publisher.publish('credit:limit_exceeded', event))
        } else {
          console.log('credit:reserved', publisher.publish('credit:reserved', event))
        }
      default:
        return
    }
 })

  app.get('/customers', async (req, res) => {
    try {
      const customers = await customerRepository.findMany()
        
      return res.json({ data: customers })
    } catch (error) {
      return res.status(500)
              .json({ message: 'Internal Server Error' })
    }
      
  })

  app.listen(PORT, () => {
    console.log('Server is listening at http://%s:%s', '0.0.0.0', PORT)  
  })
  }

start()