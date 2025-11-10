# Mock-Based Testing for Better Coverage

This document outlines how to implement mock-based unit tests to improve test coverage beyond acceptance tests.

## Why Mock-Based Testing?

- **Higher Coverage**: Test error paths and edge cases without live NetBox instance
- **Faster Execution**: No network calls or Docker setup required
- **Reliable**: Tests don't depend on external services
- **Focused**: Test specific functions in isolation

## Current Coverage Status

Current test coverage: ~75.0%

Areas that would benefit from mock tests:
- API error handling in data sources and resources
- Network failure scenarios
- Authentication errors
- Rate limiting
- Malformed responses

## Recommended Mock Testing Setup

### 1. Choose a Mocking Framework

```bash
go get github.com/stretchr/testify/mock
# or
go get github.com/golang/mock/gomock
```

### 2. Example Mock Test Structure

```go
package netbox

import (
    "errors"
    "testing"

    "github.com/fbreckle/go-netbox/netbox/client"
    "github.com/fbreckle/go-netbox/netbox/client/extras"
    "github.com/fbreckle/go-netbox/netbox/models"
    "github.com/stretchr/testify/assert"
)

// Mock client for testing
type mockNetBoxClient struct {
    extrasAPI *mockExtrasAPI
}

type mockExtrasAPI struct {
    tagsListFunc func(*extras.ExtrasTagsListParams) (*extras.ExtrasTagsListOK, error)
}

func (m *mockExtrasAPI) ExtrasTagsList(params *extras.ExtrasTagsListParams, authInfo interface{}) (*extras.ExtrasTagsListOK, error) {
    if m.tagsListFunc != nil {
        return m.tagsListFunc(params)
    }
    return nil, errors.New("mock not implemented")
}

func TestFindTag_APIError(t *testing.T) {
    // Setup mock
    mockAPI := &mockExtrasAPI{
        tagsListFunc: func(params *extras.ExtrasTagsListParams) (*extras.ExtrasTagsListOK, error) {
            return nil, errors.New("connection refused")
        },
    }

    mockClient := &client.NetBoxAPI{}
    // Note: In practice, you'd need to properly inject the mock

    // This is a simplified example - actual implementation would require
    // interface extraction and dependency injection
    tag, err := findTag(mockClient, "test-tag")

    assert.Error(t, err)
    assert.Nil(t, tag)
    assert.Contains(t, err.Error(), "API Error")
}
```

### 3. Implementation Strategy

1. **Extract Interfaces**: Create interfaces for API clients to enable mocking
2. **Dependency Injection**: Modify functions to accept interfaces instead of concrete types
3. **Mock Generation**: Use code generation tools to create mocks
4. **Test Organization**: Separate unit tests from acceptance tests

### 4. Functions to Mock Test

Priority order for mock testing:

1. **Utility Functions** (already done)
   - `findTag` in `tags.go`
   - `getNestedTagListFromResourceDataSet`
   - `readTags`

2. **Data Source Functions**
   - `dataSourceNetboxAsnsRead` - API errors, no results
   - `dataSourceNetboxTagRead` - multiple results, API errors

3. **Resource Functions**
   - CRUD operations with API failures
   - 404 handling in read/delete
   - Validation errors

4. **Provider Functions**
   - `providerConfigure` with invalid credentials
   - Version checking failures

### 5. Benefits Expected

- Coverage increase: 74.8% → 85%+
- Faster test execution
- Better error path testing
- Reduced CI resource usage

## Getting Started

1. Start with simple functions like `findTag`
2. Extract interfaces for API clients
3. Use testify/mock for simple mocking
4. Gradually expand to more complex scenarios

## Example Implementation

See `tags_mock_test.go` for a basic mock test example (requires interface extraction for full implementation).
