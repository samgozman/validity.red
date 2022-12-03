import type { Component } from "vue";
import {
  DocumentOutline,
  IdCardOutline,
  TimerOutline,
  SkullOutline,
  ApertureOutline,
  MedkitOutline,
  HomeOutline,
  CarSportOutline,
  BodyOutline,
  EarthOutline,
  SchoolOutline,
  BusinessOutline,
  RibbonOutline,
  CardOutline,
  ReaderOutline,
  QrCodeOutline,
  ReceiptOutline,
  TicketOutline,
  BusOutline,
} from "@vicons/ionicons5";

interface IDocumentType {
  name: string;
  icon: Component;
}

export class DocumentType {
  /** Document types map */
  public static types = new Map<number, IDocumentType>([
    [0, { name: "Document", icon: DocumentOutline }],
    [1, { name: "Passport", icon: IdCardOutline }],
    [2, { name: "Internal passport", icon: IdCardOutline }],
    [3, { name: "Foreign passport", icon: IdCardOutline }],
    [4, { name: "Identity card", icon: IdCardOutline }],
    [5, { name: "Driving license", icon: TimerOutline }],
    [6, { name: "Hunting license", icon: SkullOutline }],
    [7, { name: "Firearms license", icon: ApertureOutline }],
    [8, { name: "Medical insurance", icon: MedkitOutline }],
    [9, { name: "Property insurance", icon: HomeOutline }],
    [10, { name: "Vehicle insurance", icon: CarSportOutline }],
    [11, { name: "Personal insurance", icon: BodyOutline }],
    [12, { name: "Visa", icon: EarthOutline }],
    [13, { name: "Student visa", icon: SchoolOutline }],
    [14, { name: "Work permit", icon: BusinessOutline }],
    [15, { name: "Residence permit", icon: RibbonOutline }],
    [16, { name: "Credit card", icon: CardOutline }],
    [17, { name: "Certificate", icon: ReaderOutline }],
    [18, { name: "Vaccination Certificate", icon: QrCodeOutline }],
    [19, { name: "Warranty Certificate", icon: ReceiptOutline }],
    [20, { name: "Coupon", icon: TicketOutline }],
    [21, { name: "Travel card", icon: BusOutline }],
    [255, { name: "Other", icon: DocumentOutline }],
  ]);

  public static getName(typeId = 0): string {
    const t = this.types.get(typeId);
    return t ? t.name : "Unknown";
  }

  public static getIcon(typeId = 0): Component {
    const t = this.types.get(typeId);
    return t ? t.icon : DocumentOutline;
  }
}
