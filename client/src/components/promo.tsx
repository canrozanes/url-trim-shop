import React from "react"
import styled from "styled-components"
import Grid from "@material-ui/core/Grid"
import Typography from "@material-ui/core/Typography"
import logo from "../typing.jpeg"
import Box from "@material-ui/core/Box"

const StyledImage = styled.img`
  max-width: 100%;
  display: block;
`

const Promo = () => {
  return (
    <Box m="40px 0">
      <Grid container spacing={3}>
        <Grid item xs={12} md={8}>
          <Typography variant="h2" component="h1">
            Shorten your links
          </Typography>
          <Typography variant="h6" component="h2">
            This humble URL shortener does what you expect any url shortener to
            do.
          </Typography>
          <Typography variant="h6" component="h2" paragraph>
            It shortens urls.
          </Typography>
          <Typography variant="body1" component="h2" paragraph>
            This application was built by{" "}
            <a href="canrozanes.com">Can Rozanes</a> to practice building an
            HTTP Server on Golang. You can find the repo for the project{" "}
            <a href="https://github.com/canrozanes/url-trim-shop">here</a>.
          </Typography>
          <Typography variant="body1" component="h2" paragraph>
            If you have any feedback, comments or suggestions, please do not
            hesitate to reach out!
          </Typography>
        </Grid>
        <Grid item xs={false} md={4}>
          <StyledImage
            src={logo}
            alt={"closeup of a hands typing on a laptop"}
          />
        </Grid>
      </Grid>
    </Box>
  )
}

export default Promo
