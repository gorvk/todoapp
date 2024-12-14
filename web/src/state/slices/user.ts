import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { Profile } from "../../models/types";

const initialState: Profile | null = null as Profile | null ;
export const slice = createSlice({
    name: "currentUser",
    initialState,
    reducers: {
        setCurrentUser: (_, action: PayloadAction<Profile>) => {
            return action.payload;
        },
        removeCurrentUser: () => {
            return null;
        }
    }
});