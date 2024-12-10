interface RecipientAddressRequest {
  postcode: number;
  country: string;
  region: string;
  city: string;
  district: string;
  street: string;
  houseNumber: number;
  apartmentNumber: number;
}

interface RecipientAddressResponse {
  apartmentNumber: string;
  city: string;
  country: string;
  countryside: boolean;
  posteRestante: boolean;
  created: string;
  description: string;
  detailedInfo: string;
  district: string;
  foreignStreetHouseApartment: string;
  houseNumber: string;
  id: number;
  lastModified: string;
  mailbox: string;
  postcode: string;
  region: string;
  specialDestination: string;
  street: string;
}

interface RecipientObjectRequest {
  firstName: string;
  lastName: string;
  addressId: number;
  phoneNumber: number;
  type: "INDIVIDUAL" | "COMPANY" | "PRIVATE_ENTREPRENEUR";
}

interface Address {
  apartmentNumber: string;
  city: string;
  country: string;
  countryside: boolean;
  posteRestante: boolean;
  created: string;
  description: string;
  detailedInfo: string;
  district: string;
  foreignStreetHouseApartment: string;
  houseNumber: string;
  id: number;
  lastModified: string;
  mailbox: string;
  postcode: string;
  region: string;
  specialDestination: string;
  street: string;
}

interface Email {
  email: string;
  main: boolean;
  uuid: string;
}

interface Phone {
  main: boolean;
  phoneId: number;
  phoneNumber: string;
  type: "PERSONAL" | "WORK";
  uuid: string;
}

interface PostPayRecipient {
  bankAccount: string;
  counterpartyUuid: string;
  edrpou: string;
  name: string;
  personalDataApproved: boolean;
  postPayPaymentType: "POSTPAY_PAYMENT_CASH_ONLY" | "POSTPAY_PAYMENT_CARD_ONLY";
  tin: string;
  type: "INDIVIDUAL" | "LEGAL";
  uuid: string;
}

interface AddressItem {
  address: Address;
  addressId: number;
  main: boolean;
  type: "LEGAL" | "PERSONAL";
  uuid: string;
}

interface RecipientObjectResponse {
  GDPRAccept: boolean;
  GDPRRead: boolean;
  addressId: number;
  addresses: AddressItem[];
  bankAccount: string;
  checkOnDeliveryAllowed: boolean;
  contactPersonName: string;
  counterpartyUuid: string;
  documents: Record<string, any>;
  edrpou: string;
  email: string;
  emails: Email[];
  externalId: string;
  firebaseId: string;
  firstName: string;
  lastName: string;
  latinName: string;
  middleName: string;
  name: string;
  personalDataApproved: boolean;
  phoneNumber: string;
  phones: Phone[];
  postId: string;
  postPayPaymentType: "POSTPAY_PAYMENT_CASH_ONLY" | "POSTPAY_PAYMENT_CARD_ONLY";
  postPayRecipients: PostPayRecipient[];
  resident: boolean;
  tin: string;
  type: "INDIVIDUAL" | "LEGAL";
  uniqueRegistrationNumber: string;
  uuid: string;
}

const ukrPostRoutes = {};
