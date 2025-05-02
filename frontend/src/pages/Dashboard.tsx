import { Box } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { Navbar } from "../components/Navbar";
import { NoteModel } from "../types/notes";
import { NoteList } from "../components/NoteList";
import { GetNotes } from "../api/home/GetNotes";
import { useSnackbar } from "notistack";
import { ServiceErrorResponse } from "../types/errors";

export const Dashboard: React.FC = () => {
  const auth = useContext(AuthContext);
  const { enqueueSnackbar } = useSnackbar();
  if (!auth) throw new Error("AuthContext not provided");

  const [notes, setNotes] = useState<NoteModel[]>([])

  useEffect(() => {
    const fetchNotes = async () => {
      try {
        const fetchedNotes = ((await GetNotes(10, 0)) as NoteModel[]).map(note => ({
          ...note,
          created_at: new Date(note.created_at),
          updated_at: new Date(note.updated_at),
        }));
        setNotes(fetchedNotes)
      } catch (err: any) {
        let message = "";
        if (err.response && err.response.data) {
          const svcErr = err.response.data as ServiceErrorResponse;
          message = svcErr.message;
        } else {
          message = "Неверные данные или сетевая ошибка";
        }
        enqueueSnackbar(message, { variant: "error" });
      }
    }
    fetchNotes();
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        width: "100vw",
        height: "100vh",
      }}
    >
      <Navbar />
      <NoteList notes={notes} />
    </Box>
  );
};
