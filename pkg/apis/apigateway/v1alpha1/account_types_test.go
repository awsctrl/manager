package v1alpha1

import (
	"testing"

	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestStorageAccount(t *testing.T) {
	key := types.NamespacedName{
		Name:      "foo",
		Namespace: "default",
	}
	created := &Account{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		}}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &Account{}
	g.Expect(c.Create(context.TODO(), created)).To(gomega.Succeed())

	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.Succeed())
	g.Expect(fetched).To(gomega.Equal(created))

	// Test Updating the Labels
	updated := fetched.DeepCopy()
	updated.Labels = map[string]string{"hello": "world"}
	g.Expect(c.Update(context.TODO(), updated)).To(gomega.Succeed())

	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.Succeed())
	g.Expect(fetched).To(gomega.Equal(updated))

	// Test Delete
	g.Expect(c.Delete(context.TODO(), fetched)).To(gomega.Succeed())
	g.Expect(c.Get(context.TODO(), key, fetched)).ToNot(gomega.Succeed())
}
