import { Box, Card, CardContent, Typography } from "@mui/material";
import React from "react";
import { NoteModel } from "../types/notes"

type NoteProps = {
  note: NoteModel;
};

export const Note: React.FC<NoteProps> = ({ note }) => {
  return (       
  <Card sx={{ width: '100%', boxShadow: 3 }}>
    <CardContent>
      <Typography variant="h6" component="div">
        {note.title}
      </Typography>
      <Typography variant="subtitle1" color="text.secondary" gutterBottom>
        Автор: {note.username}
      </Typography>
      <Typography variant="body1" paragraph>
        {note.description}
      </Typography>
      <Box sx={{ display: 'flex', flexDirection: "column", mt: 2 }}>
        <Typography variant="caption" color="text.secondary">
          Создано: {note.created_at.toLocaleString("ru-RU")}
        </Typography>
        <Typography variant="caption" color="text.secondary">
          Обновлено: {note.updated_at.toLocaleString("ru-RU")}
        </Typography>
      </Box>
    </CardContent>
  </Card>
)};
