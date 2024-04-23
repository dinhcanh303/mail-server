import { TestSendMailConnection } from '@/types/mail.type'
import { http, httpPrivate } from '@/utils/http'

export const testSendMail = (req: TestSendMailConnection) => http.post<object>('mails/test', req)
