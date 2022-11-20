import React from "react";
import { GoogleMap, LoadScript, Marker } from "@react-google-maps/api";

const defaultMapConfig = {
  gestureHandling: "greedy",
  options: {
    fullscreenControl: false
  },
  mapContainerStyle: {
    height: "40vh",
    width: "100%"
  }
}; 

export default function GMap() {
  return (
    <>
      <LoadScript
        googleMapsApiKey={process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY}
      >
        <GoogleMap {...defaultMapConfig} center={{ lat: 0, lng: 0 }} zoom={3}>
          <Marker position={{ lat: 0, lng: 0 }} draggable={true} />

        </GoogleMap>
      </LoadScript>
    </>
  );
}
