import { configureStore } from "@reduxjs/toolkit";
import { slice as userSlice } from "./slices/user";

export const store = configureStore({
    reducer: {
        currentUser: userSlice.reducer,
    }
})

export type AppState = ReturnType<typeof store.getState>