import React from "react"
import { Toolbar, Typography, Box } from "@material-ui/core"
import Container from "@material-ui/core/Container"
import { styled } from "@material-ui/core/styles"

const StyledFooter = styled("footer")(({ theme }) => ({
  backgroundColor: theme.palette.primary.main,
  color: theme.palette.primary.contrastText,
  width: "100%",
  position: "absolute",
  bottom: 0,
}))

const Footer = () => (
  <>
    <StyledFooter>
      <Toolbar>
        <Container maxWidth="xl">
          <Typography variant="body1">
            Built by Can Rozanes using Go + React
          </Typography>
        </Container>
      </Toolbar>
    </StyledFooter>
    <Box height="64px" />
  </>
)

export default Footer
