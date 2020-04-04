
export type SignData = {
  mail: string;
  pass: string;
}

export type OAuth2CodeRedirect = {
  state?: string;
  code?: string;
}

export type StateEnum = 'open' | 'sold' | 'draft' | 'unpaid' | 'sent' | 'complete' | 'deleted'
export type ProductDetailRequest = {
  title: string;
  detail: string;
  price: number;
  state: StateEnum;
}

export type ContractDetailsResponse = {
  item_id: string;
  title: string;
  detail: string;
  price: number;
  seller_account_id: string;
  state: StateEnum;
}

export type DisplayLimit = {
  limit?: number;
  page?: number;
}

export type StateDisplayLimit = DisplayLimit & {
  state?: StateEnum;
}

export type Balance = {
  balance: number
}

export type ErrorResponseJson = {
  message: string;
}
