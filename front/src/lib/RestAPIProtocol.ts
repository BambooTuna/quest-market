
export type SignData = {
  mail: string;
  pass: string;
}

export type OAuth2CodeRedirect = {
  state?: string;
  code?: string;
}

export type StateEnum = 'open' | 'draft' | 'closed'
export type ProductDetailRequest = {
  title: string;
  detail: string;
  price: number;
  state: StateEnum;
}

export type ProductDetailResponse = {
  id: string;
  productTitle: string;
  productDetail: string;
  requestPrice: number;
  presenterId: string;
  state: StateEnum;
}

export type DisplayLimit = {
  limit?: number;
  page?: number;
}

export type StateDisplayLimit = DisplayLimit & {
  state?: StateEnum;
}

export type ErrorResponseJson = {
  message: string;
}
