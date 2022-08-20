import type { RouteLocationRaw } from "vue-router";

export interface NavbarItem {
  name: string;
  url: RouteLocationRaw;
  iconClass: string;
}
