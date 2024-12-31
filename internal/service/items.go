package service

import (
	"github.com/google/uuid"
	"github.com/ursulgwopp/azamon/internal/errors"
	"github.com/ursulgwopp/azamon/internal/models"
)

func validateName(name string) error {
	if len(name) < 1 || len(name) > 200 {
		return errors.ErrInvalidName
	}

	return nil
}

func validateDescription(description string) error {
	if len(description) < 1 || len(description) > 2000 {
		return errors.ErrInvalidDescription
	}

	return nil
}

func validateQuantity(quantity int) error {
	if quantity < 0 {
		return errors.ErrInvalidQuantity
	}

	return nil
}

func validatePrice(price float64) error {
	if price < 0 {
		return errors.ErrInvalidPrice
	}

	return nil
}

func validateSeller(s *Service, username string, itemId uuid.UUID) error {
	seller, err := s.repo.CheckItemSeller(itemId)
	if err != nil {
		return err
	}

	if username != seller {
		return errors.ErrAccessToItemDenied
	}

	return nil
}

func (s *Service) GetItem(itemId uuid.UUID) (models.Item, error) {
	if err := s.repo.CheckItemIdExists(itemId); err != nil {
		return models.Item{}, err
	}

	return s.repo.GetItem(itemId)
}

func (s *Service) ListItems(username string) ([]models.Item, error) {
	if err := s.repo.CheckUsernameExists(username); err != nil {
		return []models.Item{}, err
	}

	return s.repo.ListItems(username)
}

func (s *Service) CreateItem(username string, req models.ItemRequest) (models.Item, error) {
	if err := validateName(req.Name); err != nil {
		return models.Item{}, err
	}

	if err := validateDescription(req.Description); err != nil {
		return models.Item{}, err
	}

	if err := validateQuantity(req.Quantity); err != nil {
		return models.Item{}, err
	}

	if err := validatePrice(req.Price); err != nil {
		return models.Item{}, err
	}

	return s.repo.CreateItem(username, req)
}

func (s *Service) UpdateItem(username string, itemId uuid.UUID, req models.ItemRequest) (models.Item, error) {
	if err := validateSeller(s, username, itemId); err != nil {
		return models.Item{}, err
	}

	if err := s.repo.CheckItemIdExists(itemId); err != nil {
		return models.Item{}, err
	}

	if err := validateName(req.Name); err != nil {
		return models.Item{}, err
	}

	if err := validateDescription(req.Description); err != nil {
		return models.Item{}, err
	}

	if err := validateQuantity(req.Quantity); err != nil {
		return models.Item{}, err
	}

	if err := validatePrice(req.Price); err != nil {
		return models.Item{}, err
	}

	return s.repo.UpdateItem(itemId, req)
}

func (s *Service) DeleteItem(username string, itemId uuid.UUID) error {
	if err := validateSeller(s, username, itemId); err != nil {
		return err
	}

	if err := s.repo.CheckItemIdExists(itemId); err != nil {
		return err
	}

	return s.repo.DeleteItem(itemId)
}
