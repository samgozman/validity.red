import type { RouteLocationRaw } from "vue-router";
import type { Component } from "vue";

export interface NavbarItem {
  name: string;
  url: RouteLocationRaw;
  icon: Component;
}
