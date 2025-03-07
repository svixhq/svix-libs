// This file is @generated
package com.svix.models;

import com.fasterxml.jackson.annotation.JsonAutoDetect;
import com.fasterxml.jackson.annotation.JsonAutoDetect.Visibility;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.svix.Utils;

import lombok.EqualsAndHashCode;
import lombok.ToString;

@ToString
@EqualsAndHashCode
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonAutoDetect(getterVisibility = Visibility.NONE, setterVisibility = Visibility.NONE)
public class GithubConfigOut {
    public GithubConfigOut() {}

    /**
     * Create an instance of GithubConfigOut given an JSON string
     *
     * @param jsonString JSON string
     * @return An instance of GithubConfigOut
     * @throws JsonProcessingException if the JSON string is invalid with respect to GithubConfigOut
     */
    public static GithubConfigOut fromJson(String jsonString) throws JsonProcessingException {
        return Utils.getObjectMapper().readValue(jsonString, GithubConfigOut.class);
    }

    /**
     * Convert an instance of GithubConfigOut to an JSON string
     *
     * @return JSON string
     */
    public String toJson() throws JsonProcessingException {
        return Utils.getObjectMapper().writeValueAsString(this);
    }
}
