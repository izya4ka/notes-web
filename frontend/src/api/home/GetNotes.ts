import { NoteModel } from "../../types/notes";
import { ServiceErrorResponse } from "../../types/errors";
import { api } from "../axios";

export const GetNotes = (amount: number, page: number): Promise<NoteModel[] | ServiceErrorResponse> => {
  return api
    .get<NoteModel[] | ServiceErrorResponse>("/notes", { params: {
      amount,
      page
    } })
    .then(res => res.data)
};
