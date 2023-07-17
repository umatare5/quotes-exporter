// (C) 2023 by Marco Paganini <paganini@paganini.net>
//
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package twelvedata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	twelvedataURL = "https://api.twelvedata.com/price?symbol=%s&apikey=%s"
)

type price struct {
	Price string `json:"price"`
}

// Quote returns the current value of a symbol.
func Quote(symbol string, apikey string) (float64, error) {
	symbol = strings.ToUpper(symbol)

	resp, err := http.Get(fmt.Sprintf(twelvedataURL, symbol, apikey))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data price
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return 0, err
	}

	if data.Price == "" {
		fmt.Println("Price is not included in JSON:", err)
		return 0, err
	}

	price, err := strconv.ParseFloat(data.Price, 64)
	if err != nil {
		fmt.Println("Error when type conversion string to float64:", err)
		return 0, err
	}

	return price, nil
}
