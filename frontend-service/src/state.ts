import { ref } from "vue";

interface IUserState {
  isAuthenticated: boolean;
  timezone: string;
  calendarId: string;
  calendarCurrentDate: string;
}

export const state = ref({
  user: {
    isAuthenticated: false,
    timezone: Intl.DateTimeFormat().resolvedOptions().timeZone,
    calendarId: "",
  } as IUserState,
});

export const setUser = (user: IUserState) => {
  state.value.user = user;
  localStorage.setItem("user", JSON.stringify(state.value.user));
};

export const setIsAuthenticated = (value: boolean) => {
  state.value.user.isAuthenticated = value;
  localStorage.setItem("user", JSON.stringify(state.value.user));
};

export const setUsersTimezone = (value: string) => {
  state.value.user.timezone = value;
  localStorage.setItem("user", JSON.stringify(state.value.user));
};

export const setCalendarId = (value: string) => {
  state.value.user.calendarId = value;
  localStorage.setItem("user", JSON.stringify(state.value.user));
};

export const setCalendarCurrentDate = (value: string) => {
  state.value.user.calendarCurrentDate = value;
  localStorage.setItem("user", JSON.stringify(state.value.user));
};
