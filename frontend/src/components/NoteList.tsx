import React from "react";
import { NoteModel } from "../types/notes";
import { Container, List, ListItem, Typography } from "@mui/material";
import { Note } from "./Note";

type NoteListProps = {
  notes: NoteModel[] | undefined;
};

export const NoteList: React.FC<NoteListProps> = ({ notes }) => {

  if (!notes || notes.length === 0) {
    return (
      <Container sx={{ height:"100%", mt: 4 }}>
        <Typography variant="h6" color="textSecondary">
          Заметки не найдены
        </Typography>
      </Container>
    );
  }

  return (
    <Container sx={{ mt: 4, mb: 4, display: "flex", height: "100%", width: "100%", flexDirection: "column", alignItems: "flex-start" }}>
      <List>
        {notes.map((note) => (
          <ListItem key={note.id} sx={{ mb: 2 }}>
            <Note note={note} />
          </ListItem>
        ))}
      </List>
    </Container>
  );
};